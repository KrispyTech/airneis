package auth

import (
	"fmt"
	"time"

	"github.com/KrispyTech/airneis/lib/shared/airror"
	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/lib/shared/vault"
	"github.com/KrispyTech/airneis/src/serializers"
	"github.com/golang-jwt/jwt/v5"
)

func getJWTParameters(vlt vault.VaultClient) (string, int, error) {
	secretKey, err := vlt.ReadSecret(constants.JWT)
	if err != nil {
		return "", 0, airror.New(err, "Vault Error: Unable to get JWT")
	}

	expirationIn, err := vlt.ReadSecret(constants.JWTExpiration)
	if err != nil {
		return "", 0, airror.New(err, "Vault Error: Unable to get JWT Expiration")
	}

	jwtExpiration, err := helpers.ConvertStringToInt(expirationIn)
	if err != nil {
		return "", 0, airror.New(err, "Vault Error: Unable convert to string expiration duration")
	}

	return secretKey, jwtExpiration, nil
}

func CreateToken(user serializers.SerializedUser, vlt vault.VaultClient) (string, error) {
	rawToken, expirationTime, err := getJWTParameters(vlt)
	if err != nil {
		return "", err
	}

	jwtPayload := map[string]serializers.SerializedUser{
		"userData": user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"payload": jwtPayload,
		"exp":     time.Now().Add(time.Minute * time.Duration(expirationTime)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(rawToken))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckToken(token string, vlt vault.VaultClient) (serializers.SerializedUser, error) {
	rawToken, _, err := getJWTParameters(vlt)
	if err != nil {
		return serializers.SerializedUser{}, airror.New(err, "Failed to get JWT params")
	}

	parsedToken, err := jwt.Parse(token, func(_ *jwt.Token) (interface{}, error) { return []byte(rawToken), nil })
	if err != nil {
		return serializers.SerializedUser{}, airror.New(err, "Unable to parse token")
	}

	if !parsedToken.Valid {
		return serializers.SerializedUser{}, airror.New(err, "invalid token")
	}

	expirationTime, err := parsedToken.Claims.GetExpirationTime()
	if err != nil {
		return serializers.SerializedUser{}, airror.New(err, "unable to get expiration time")
	}

	if expirationTime.Before(time.Now()) {
		return serializers.SerializedUser{}, fmt.Errorf("token has expired")
	}

	return serializers.SerializedUser{
		FirstName: parsedToken.Claims.(jwt.MapClaims)["payload"].(map[string]interface{})["userData"].(map[string]interface{})["firstName"].(string),
		LastName:  parsedToken.Claims.(jwt.MapClaims)["payload"].(map[string]interface{})["userData"].(map[string]interface{})["lastName"].(string),
		Email:     parsedToken.Claims.(jwt.MapClaims)["payload"].(map[string]interface{})["userData"].(map[string]interface{})["email"].(string),
		IsActive:  parsedToken.Claims.(jwt.MapClaims)["payload"].(map[string]interface{})["userData"].(map[string]interface{})["isActive"].(bool),
		IsAdmin:   parsedToken.Claims.(jwt.MapClaims)["payload"].(map[string]interface{})["userData"].(map[string]interface{})["isAdmin"].(bool),
	}, nil
}

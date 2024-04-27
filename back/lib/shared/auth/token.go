package auth

import (
	"fmt"
	"time"

	"github.com/KrispyTech/airneis/lib/shared/vault"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/src/serializers"
	"github.com/golang-jwt/jwt/v5"
)

func getJWTParameters(vlt vault.VaultClient) (string, int, error) {
	secretKey, err := vlt.ReadSecret(constants.JWT)
	if err != nil {
		return "", 0, fmt.Errorf("Vault Error: Unable to get JWT. %s", err.Error())
	}

	expirationIn, err := vlt.ReadSecret(constants.JWTExpiration)
	if err != nil {
		return "", 0, fmt.Errorf("Vault Error: Unable to get JWT Expiration. %s", err.Error())
	}

	jwtExpiration, err := helpers.ConvertStringToInt(expirationIn)
	if err != nil {
		return "", 0, fmt.Errorf("Vault Error: Unable convert to string expiration duration. %s", err.Error())
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

	tokenString, err := token.SignedString(rawToken)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckToken(token string, vlt vault.VaultClient) (serializers.SerializedUser, error) {
	rawToken, _, err := getJWTParameters(vlt)
	if err != nil {
		return serializers.SerializedUser{}, err
	}

	parsedToken, err := jwt.Parse(token, func(_ *jwt.Token) (interface{}, error) {
		return rawToken, nil
	})

	if err != nil {
		return serializers.SerializedUser{}, fmt.Errorf("Unable to parse token. %s", err.Error())
	}

	if !parsedToken.Valid {
		return serializers.SerializedUser{}, fmt.Errorf("invalid token")
	}

	return getUserDataFromParsedToken(parsedToken), nil
}

func getUserDataFromParsedToken(parsedToken *jwt.Token) serializers.SerializedUser {
	var jwtPayload = parsedToken.Claims.(jwt.MapClaims)["payload"]

	return jwtPayload.(map[string]interface{})["userData"].(serializers.SerializedUser)
}

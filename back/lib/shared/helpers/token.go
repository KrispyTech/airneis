package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/src/serializers"
	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey   = []byte(os.Getenv(constants.EnvJWTSecretKey))
	expireIn, _ = ConvertStringToInt(os.Getenv(constants.EnvJWTExpire))
)

func CreateToken(user serializers.SerializedUser) (string, error) {
	jwtPayload := map[string]serializers.SerializedUser{
		"userData": user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"payload": jwtPayload,
		"exp":     time.Now().Add(time.Minute * time.Duration(expireIn)).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckToken(token string) (serializers.SerializedUser, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return serializers.SerializedUser{}, err
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

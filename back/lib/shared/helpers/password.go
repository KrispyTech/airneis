package helpers

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"math/rand"

	"github.com/gofiber/fiber/v2"

	"github.com/KrispyTech/airneis/lib/shared/constants"
)

func GenerateSalt() string {
	salt := make([]byte, constants.SaltLength)

	for i := range salt {
		salt[i] = constants.Chars[rand.Intn(len(constants.Chars))]
	}

	return string(salt)
}

func CreateHashedPasswordAndSalt(password string) (string, string, error) {
	salt := GenerateSalt()
	sha := sha512.New()
	if _, err := sha.Write([]byte(password + salt)); err != nil {
		return "", "", err
	}
	hashedPassword := hex.EncodeToString(sha.Sum(nil))

	return hashedPassword, salt, nil
}

func HashPasswordWithSalt(password string, salt string) (string, error) {
	sha := sha512.New()
	if _, err := sha.Write([]byte(password + salt)); err != nil {
		return "", err
	}
	hashedPassword := hex.EncodeToString(sha.Sum(nil))

	return hashedPassword, nil
}

type inputData struct {
	Password string `json:"password"`
}

func GetPasswordFromContext(ctx *fiber.Ctx) (string, error) {
	var inputData inputData

	if err := json.Unmarshal(ctx.Body(), &inputData); err != nil {
		return "", err
	}

	password := inputData.Password

	return password, nil
}

func GetHashedPasswordAndSaltFromContext(ctx *fiber.Ctx) (string, string, error) {
	password, err := GetPasswordFromContext(ctx)

	if err != nil {
		return "", "", err
	}

	hashedPassword, salt, err := CreateHashedPasswordAndSalt(password)

	if err != nil {
		return "", "", err
	}

	return hashedPassword, salt, nil
}

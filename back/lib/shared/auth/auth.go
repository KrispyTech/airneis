package auth

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"

	"golang.org/x/crypto/scrypt"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
)

func CreatePassword(authData helpers.AuthData) (hashedPassword string, salt string, err error) {
	hashedPassword, salt, err = createNewPasswordHash(authData.Password)
	if err != nil {
		return "", "", err
	}

	return hashedPassword, salt, nil
}

func createNewPasswordHash(rawPassword string) (password string, salt string, err error) {
	salt, err = generateSalt()
	if err != nil {
		return "", "", err
	}

	hashedPassword, err := hashPassword([]byte(rawPassword), []byte(salt))
	if err != nil {
		return "", "", err
	}

	password = hex.EncodeToString(hashedPassword)

	return password, salt, nil
}

func generateSalt() (string, error) {
	salt := make([]byte, constants.SaltLength)
	max := big.NewInt(int64(len(constants.Chars)))

	for i := range salt {
		randomNumber, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		salt[i] = constants.Chars[randomNumber.Int64()]
	}

	println(string(salt))

	return string(salt), nil
}

func hashPassword(rawPassword, salt []byte) ([]byte, error) {
	const (
		/*It controls the amount of work required to compute the derived key and directly affects both:
		CPU time and memory usage.*/
		memoryCost = 32768

		// It represents the block size used during the mixing step of the scrypt algorithm
		memoryUsedByAlgo = 8

		//  It is used to determine the number of parallel chains of computation.
		parallelizationFactor = 1

		/*It specifies the length of the derived key in bytes.
		In cryptographic terms, it indicates the desired length of the output key. */
		keyLength = 32
	)

	hashedPassword, err := scrypt.Key(rawPassword, salt, memoryCost, memoryUsedByAlgo, parallelizationFactor, keyLength)
	if err != nil {
		return nil, err
	}

	return hashedPassword, nil
}

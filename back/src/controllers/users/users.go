package users

import (
	"fmt"

	"github.com/KrispyTech/airneis/lib/shared/airror"

	"github.com/KrispyTech/airneis/lib/shared/auth"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func fillUserInformation(ctx *fiber.Ctx, user model.User) (model.User, error) {
	authData, err := helpers.GetAuthContext(ctx)
	if err != nil {
		return model.User{}, helpers.SendInternalError(ctx, err)
	}

	if len(authData.Email) == 0 {
		return model.User{}, fmt.Errorf("please provide an email")
	}

	if len(authData.Password) == 0 {
		return model.User{}, fmt.Errorf("please provide a password")
	}

	hashedPassword, salt, err := auth.CreatePassword(authData)
	if err != nil {
		return model.User{}, helpers.SendBadRequest(ctx, err)
	}

	user.PasswordHash, user.PasswordSalt = hashedPassword, salt

	return user, nil
}

func GetUserByEmail(email string) (model.User, error) {
	var user model.User

	if len(email) == 0 {
		return model.User{}, fmt.Errorf("no email")
	}

	if err := config.Database.Where("email = ?", email).First(&user).Error; err != nil {
		return model.User{}, airror.New(err, "user not found")
	}

	return user, nil
}

func checkUser(user model.User) error {
	if len(user.FirstName) == 0 {
		return fmt.Errorf("please provide a first name")
	}

	if len(user.LastName) == 0 {
		return fmt.Errorf("please provide a last name")
	}

	if len(user.Email) == 0 {
		return fmt.Errorf("please provide an email")
	}

	return nil
}

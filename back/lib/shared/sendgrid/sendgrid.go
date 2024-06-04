package sendgrid

import (
	"fmt"
	"net/http"

	"github.com/KrispyTech/airneis/lib/shared/vault"
	"github.com/go-playground/validator"
	"github.com/pkg/errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendgridAPI interface {
	ValidateMail(email Mail) (*mail.SGMailV3, error) 
	SendMail(mail *mail.SGMailV3) (err error)
}

type SendgridClient struct {
	Api *sendgrid.Client
}

type Mail struct {
	From    User   `validate:"required"`
	To      User   `validate:"required"`
	Subject string `validate:"required"`
	Content string `validate:"required"`
}

type User struct {
	Name      string `validate:"required"`
	FirstName string
	LastName  string
	Email     string `validate:"required,email"`
}


func InitializeSendgridAPI(vc vault.VaultClient) (client SendgridAPI, err error) {
	token, err := vc.ReadSecret("SENDGRID_API_KEY")
	if err != nil {
		return nil, errors.Wrapf(err, "GetNeonCreds, unable to get server")
	}

	return &SendgridClient{
		Api: sendgrid.NewSendClient(token),
	}, nil

}

func (sg *SendgridClient) ValidateMail(email Mail) (*mail.SGMailV3, error) {
	validate := validator.New()
	from := mail.NewEmail(email.From.Name, email.From.Email)
	to := mail.NewEmail(email.To.Name, email.To.Email)
	message := mail.NewSingleEmailPlainText(from, email.Subject, to, email.Content)

	if err := validate.Struct(email); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return nil, fmt.Errorf("Validation error on field '%s': %s\n", err.StructField(), err.Tag())
		}
	}

return message, nil
}

func (sg *SendgridClient) SendMail(mail *mail.SGMailV3) (err error) {
	res, err := sg.Api.Send(mail)
	if err != nil {
		return fmt.Errorf("Unable to send Email")
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Error %d: %s", res.StatusCode, res.Body)
	}

	return nil
}
package mailer

import (
	"github.com/matcornic/hermes/v2"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"net/http"
	"os"
)

type sendMail struct {
}

type SendMailer interface {
	SendResetPassword(string, string, string, string, string) (*EmailResponse, error)
}

var (
	SendMail SendMailer = &sendMail{}
)

type EmailResponse struct {
	Status   int
	RespBody string
}

func (s *sendMail) SendResetPassword(ToUser string, FromAdmin string, Token string, SendGridKey string, AppEnv string) (*EmailResponse, error) {
	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "Handl3r",
			Link: "handl3r.com",
		},
	}
	var forgotUrl string
	if os.Getenv("APP_ENV") == "production" {
		forgotUrl = "http://handl3r.com/resetpassword/" + Token
	} else {
		forgotUrl = "http://localhost:3000/resetpassword/" + Token
	}
	email := hermes.Email{
		Body: hermes.Body{
			Name: ToUser,
			Intros: []string{
				"Welcome to handl3r! Good to see you.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Click this link to reset your password",
					Button: hermes.Button{
						Color: "#FFFFFF",
						Text:  "ResetPassword",
						Link:  forgotUrl,
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return nil, err
	}

	from := mail.NewEmail("hand3r", FromAdmin)
	subject := "ResetPassword"
	to := mail.NewEmail("Reset Password", ToUser)
	message := mail.NewSingleEmail(from, subject, to, emailBody, emailBody)
	client := sendgrid.NewSendClient(SendGridKey)
	_, err = client.Send(message)
	if err != nil {
		return nil, err
	}
	return &EmailResponse{
		Status:   http.StatusOK,
		RespBody: "Success, Please click on the link provided in your email",
	}, nil
}

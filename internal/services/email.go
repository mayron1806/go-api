package services

import (
	"fmt"
	"net/smtp"

	"github.com/mayron1806/go-api/config"
)

type EmailService struct {
	auth      *smtp.Auth
	sendEmail bool
	address   string
	from      string
	logger    *config.Logger
}

func NewEmailService() *EmailService {
	logger := config.GetLogger("Email")
	env := config.GetEnv()
	if !env.SHOULD_SEND_EMAILS {
		return &EmailService{
			logger:    logger,
			sendEmail: false,
		}
	}
	auth := smtp.PlainAuth("", env.SMTP_USER, env.SMTP_PASS, env.SMTP_HOST)
	return &EmailService{
		logger:    logger,
		auth:      &auth,
		sendEmail: true,
		address:   fmt.Sprintf("%s:%s", env.SMTP_HOST, env.SMTP_PORT),
		from:      env.SMTP_FROM,
	}
}
func (e *EmailService) SendEmail(to string, subject string, body string) error {
	if !e.sendEmail {
		e.logger.Debug("email sending is disabled")
		e.logger.Debugf("to: %s, subject: %s, body: %s", to, subject, body)
		return nil
	}
	err := smtp.SendMail(e.address, *e.auth, e.from, []string{to}, []byte(body))
	if err != nil {
		e.logger.Errorf("error sending email: %s", err.Error())
		return err
	}
	return nil
}

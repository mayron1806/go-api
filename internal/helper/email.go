package helper

import "net/mail"

func VerifyIsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

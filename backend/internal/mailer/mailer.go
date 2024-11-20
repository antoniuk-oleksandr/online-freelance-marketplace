package mailer

import (
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmail(recipient string, subject string, body string, bodyType string) error {
 	smtpHost := "smtp.gmail.com"
    smtpPort := 587
    senderEmail := os.Getenv("SENDER_EMAIL")
    password := os.Getenv("EMAIL_PASSWORD")
    
    m := gomail.NewMessage()
    m.SetHeader("From", senderEmail)
    m.SetHeader("To", recipient)
    m.SetHeader("Subject", subject)
    m.SetBody(bodyType, body)

    d := gomail.NewDialer(smtpHost, smtpPort, senderEmail, password)
    if err := d.DialAndSend(m); err != nil {
    	return err
    }

    return nil
}
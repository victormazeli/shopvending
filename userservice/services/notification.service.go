package services

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/mailgun/mailgun-go/v3"
	"gopkg.in/gomail.v2"
	"time"
)

type NotificationService struct{}

func (n NotificationService) SendEmail(receiver, body, subject string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "cfttest728@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", receiver)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.mailtrap.io", 2525, "2cc985aa082d48", "a58a393363be2b")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func (n NotificationService) SendSimpleMessage(receiver, body, subject string) (string, error) {
	mg := mailgun.NewMailgun("cftchurchesdevenvironment.xyz", "1f321d0b6de8db68c15760834b5e10db-07e2c238-95c80d39")
	m := mg.NewMessage(
		"noreply@cftchurchesdevenvironment.xyz",
		subject,
		body,
		receiver,
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}

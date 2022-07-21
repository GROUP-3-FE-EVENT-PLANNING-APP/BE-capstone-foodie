package helper

import (
	"fmt"
	"os"

	gomail "gopkg.in/mail.v2"
)

type Recipient struct {
	Name         string
	Email        string
	Handphone    string
	TotalPayment uint64
}

// 465
// 587

func SendEmail(data Recipient) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "altacapstonegroup3@gmail.com")
	msg.SetHeader("To", "tesbahaso1503@gmail.com")
	msg.SetHeader("Subject", "Notification Payment")
	msg.SetBody("text/html", "<b>This is the body of the mail</b>")

	n := gomail.NewDialer(os.Getenv("GMAIL_SMTP_HOST"), 587, os.Getenv("GMAIL_AUTH_EMAIL"), os.Getenv("GMAIL_AUTH_PASSWORD"))

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		fmt.Println(err.Error())
	}
}

// const CONFIG_SMTP_HOST = "smtp-mail.outlook.com"
// const CONFIG_SMTP_PORT = 587

// func SendEmail(data Recipient) {
// 	mailer := gomail.NewMessage()
// 	mailer.SetHeader("From", "PT. Makmur Subur Jaya <altacapstonegroup3@outlook.com>")
// 	mailer.SetHeader("To", "tesbahaso1503@gmail.com")
// 	mailer.SetAddressHeader("Cc", "tesbahaso1503@gmail.com", "Tra Lala La")
// 	mailer.SetHeader("Subject", "Email Notif")
// 	mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")

// 	dialer := gomail.NewDialer(
// 		os.Getenv("EMAIL_SMTP_HOST"),
// 		CONFIG_SMTP_PORT,
// 		os.Getenv("EMAIL_AUTH_EMAIL"),
// 		"altaGroup3_",
// 	)

// 	err := dialer.DialAndSend(mailer)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		// log.Fatal(err.Error())
// 	}
// }

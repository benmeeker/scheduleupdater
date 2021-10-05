package main

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func notify(uploadFile string) {
	e := email.NewEmail()
	e.AttachFile(uploadFile)
	// TDOO: Update the email header
	e.From = "name <@gmail.com>"
	// TODO: Put in your (TO) email addresses below
	e.To = []string{"Email addresses"}
	e.Subject = "Schedule has been updated"
	e.Text = []byte("Weekly Schedule")
	// TODO: Put in your email and details
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "@gmail.com", "password", "smtp.gmail.com"))
	if err != nil {
		log.Print(err)
	}
}

package main

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func notify(uploadFile string) {
	e := email.NewEmail()
	e.AttachFile(uploadFile)
	e.From = "schedule updater <scheduleupdater@gmail.com>"
	e.To = []string{"ben@meekers.org", "ublaknoob@yahoo.com", "schultz.15@icloud.com", "l_gardner1991@yahoo.com", "coleshrop@gmail.com", "jonahlarsen001@gmail.com", "harrisonhthomas@gmail.com", "jacobbird.birdes@gmail.com"}
	e.Subject = "Schedule has been updated"
	e.Text = []byte("Weekly Schedule")
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "scheduleupdater@gmail.com", "schedule123", "smtp.gmail.com"))
	if err != nil {
		log.Print(err)
	}
}

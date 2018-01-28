package main

import (
	"log"
	"net/smtp"
	"strings"
)

var (
	smtpHost   = "smtp.gmail.com"
	smtpPort   = "587"
	sender     = "sender@gmail.com"
	password   = "SecretPassword"
	recipients = []string{
		"recipient1@gmail.com",
		"recipient2@gmail.com",
	}
	subject = "Subject Line"
)

func main() {
	auth := smtp.PlainAuth("", sender, password, smtpHost)

	textEmail := []byte(
		`To: ` + strings.Join(recipients, ", ") + `
Mime-Version: 1.0
Content-Type: text/plain; charset="UTF-8";
Subject: ` + subject + `

Hello,

This is a plain text email.
`)

	htmlEmail := []byte(
		`To: ` + strings.Join(recipients, ", ") + `
Mime-Version: 1.0
Content-Type: text/html; charset="UTF-8";
Subject: ` + subject + `

<html>
<h1>Hello</h1>
<hr />
<p>This is an <strong>HTML</strong> email.</p>
</html>
`)

	// Send text version of email
	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		sender,
		recipients,
		textEmail,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Send HTML version
	err = smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		sender,
		recipients,
		htmlEmail,
	)
	if err != nil {
		log.Fatal(err)
	}
}

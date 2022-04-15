package main

import (
	"fmt"
	"log"
	"net/smtp"
)

/*
const (
	USERNAME = "user"
	PASSWD   = "123"
	HOST     = "emx.pagenation.com"
)
*/

const (
	USERNAME = "pete@kucinta.com"
	PASSWD   = "We!@com3"
	HOST     = "smtppro.zoho.com"
)

func main() {
	from := "pete@kucinta.com"
	to := []string{
		"editor@pagenation.com",
		"ptong@mac-net.com",
	}
	msg := []byte(
		"From: <pete@kucinta.com>\r\n" +
			"To: <editor@pagenation.com>" +
			",<ptong@mac-net.com>\r\n" +
			"Subject: Golang testing mail 3\r\n" +
			"Welcome to Go!\r\n")

	auth := smtp.PlainAuth("", USERNAME, PASSWD, HOST)
	err := smtp.SendMail(HOST+":25", auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mail sent successfully!")
}

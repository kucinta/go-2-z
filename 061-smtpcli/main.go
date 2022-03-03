package main

import (
	"log"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

func main() {
	// Setup authentication information.
	auth := sasl.NewPlainClient("", "user", "123")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"pete@eldus.com"}
	msg := strings.NewReader("To: pete@eldus.com\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("emx.pagenation.com:587", auth, "user@emx.pagenation.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/smtp"
)

// TLS Email Example

func main() {
	fmt.Println("Run")
	send2("hello there")
	//send1()
}

func send2(body string) {
	user := "pete@kucinta.com"
	pass := "We!@com3"
	fm := "hpcalc.warranty@educalc.net"
	to := "szeye@hotmail.com"
	subj := "Test Message"

	msg := "From: " + fm + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subj + "\n\n" +
		body

	fmt.Println("Sending via 587...")

	err := smtp.SendMail("smtppro.zoho.com:587", smtp.PlainAuth("", user, pass, "smtppro.zoho.com"), user, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Print(body + " from " + fm)
	log.Print("Sent to " + to + "!")

}

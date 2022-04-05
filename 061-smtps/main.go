package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

// TLS Email Example

func main() {

	from := mail.Address{"", "ptong@mac-net.com"}
	to := mail.Address{"", "szeye@hotmail.com"}
	subj := "This is the email subject"
	body := "This is an example body.\nWith two lines.\nplus one :)"

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	println(headers["From"])
	println("--")
	println(message)
	println("--")
	println(from.String())
	message = "From: <ptong@mac-net.com>\r\n"
	message += "To: <szeye@hotmail.com>\r\n"
	message += "Subject: This is the email subject 123\r\n"
	message += body

	// Connect to the SMTPS Server
	servername := "emx.pagenation.com:465"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", "user", "123", host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	// To
	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

}

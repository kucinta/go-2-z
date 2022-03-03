package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

// SSL/TLS Email Example

func main() {

	from := mail.Address{"", "user@emx.pagenation.com"}
	to := mail.Address{"", "ptong@mac-net.com"}
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
	message = "From: <user@emx.pagenation.com>\r\n"
	message += "To: <ptong@mac-net.com>\r\n"
	message += "Subject: This is the email subject 123\r\n"
	message += body

	// Connect to the SMTP Server
	servername := "emx.pagenation.com:465" //465
	//servername := "192.168.1.66:25"

	host, _, _ := net.SplitHostPort(servername)

	println(servername)
	println(host)

	auth := smtp.PlainAuth("", "user", "123", host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
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

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

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

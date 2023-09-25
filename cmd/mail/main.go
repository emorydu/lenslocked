package main

import (
	"fmt"
	"github.com/go-mail/mail/v2"
)

func main() {
	const (
		host     = "sandbox.smtp.mailtrap.io"
		port     = 587
		username = "952c04fa231288"
		password = "bcc3ef043729de"
	)

	from := "orangeduxiaocheng@gmail.com"
	to := "2392158146@qq.com"
	subject := "This is test email"
	plaintext := "This is the body of the email"
	html := `<h1>Hello there buddy!</h1><p>This is the email</p><p>Hope you enjoy it</p>`

	msg := mail.NewMessage()
	msg.SetHeader("To", to)
	msg.SetHeader("From", from)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", plaintext)
	msg.AddAlternative("text/html", html)
	//msg.WriteTo(os.Stdout)

	dialer := mail.NewDialer(host, port, username, password)
	err := dialer.DialAndSend(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Message sent")
}

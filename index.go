package main

import (
	"fmt"
	"net/smtp"
)

func sendMail(subject string, body string, to []string) {
	atuth := smtp.PlainAuth(

		//  prepaer the email auth service
		"",
		"fromemail@gmail.com",
		"here you need the password you get after allowng two factor auth if you use gmail",
		"smtp.gmail.com",
	)
	msg := "Subject:" + subject + "\n" + body
	// send email
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		atuth,
		// Sendign from
		"fromemail@gmail.com",

		// sending email to , this is slice
		to,
		//  message or the content of the email
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	sendMail("My Daily Report", "HEllo Mr Pima Please find attached for the daily report", []string{"receiver@gmail.com"})
	fmt.Println("Helo")
}

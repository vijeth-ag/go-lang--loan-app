package services

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("err loadung .env", err)
	}
	fromEmail := os.Getenv("FROM_EMAIL_ADDRESS")

	log.Println("fromEmail", fromEmail)

}

func SendMail(subject string, message string, to string) bool {

	log.Println("to", to)

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("err loadung .env", err)
	}
	fromEmail := os.Getenv("FROM_EMAIL_ADDRESS")

	log.Println("fromEmail", fromEmail)

	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)

	m.SetHeader("To", to)

	m.SetHeader("Subject", subject)

	m.SetBody("text/html", message)

	log.Println("fromEmail", fromEmail)
	log.Println("to", to)

	d := gomail.NewDialer("smtp.gmail.com", 587, fromEmail, os.Getenv("EMAIL_APP_PASSWORD"))

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
	return true
}

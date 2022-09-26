package services

import (
	"auth/db"
	"auth/utils"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	gomail "gopkg.in/mail.v2"
)

func SendActivationMail(to string) error {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("err loadung .env")
	}
	fromEmail := os.Getenv("ACTIVATION_FROM_EMAIL_ADDRESS")
	activationUrl := os.Getenv("ACTIVATION_URL")

	code := utils.RandomString(6)

	db.Set(to, code)

	activationUrl = activationUrl + "?email=" + to + "&code=" + code

	activationLink := "<a href=\"$activationLink$\">Activate account</a>"

	activationLink = strings.Replace(activationLink, "$activationLink$", activationUrl, -1)

	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)

	m.SetHeader("To", to)

	m.SetHeader("Subject", "Gomail test subject")

	m.SetBody("text/html", activationLink)

	log.Println("fromEmail", fromEmail)
	log.Println("to", to)

	d := gomail.NewDialer("smtp.gmail.com", 587, fromEmail, "lymcritkpqqpdnlo")

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
	return nil
}

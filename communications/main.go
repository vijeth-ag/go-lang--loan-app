package main

import (
	rabbitmq "communications/rabbitMq"
	"communications/services"
	"encoding/json"
	"fmt"
	"log"
)

type MessageBody struct {
	UserEmail string
	Message   string
	DateTime  string
}

func main() {

	rabbitMqConn, err := rabbitmq.Connect()
	if err != nil {
		log.Println("err getting rabbitMq conn")
	}

	ch, err := rabbitMqConn.Channel()

	if err != nil {
		log.Println("err getting rabbitMqConn channel", err)
	}

	msgs, err := ch.Consume(
		"UserCommunications",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println("err initing Consume", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Rcvd msg: %s\n", d.Body)

			messageBody := MessageBody{}
			err := json.Unmarshal(d.Body, &messageBody)
			if err != nil {
				log.Println("err unmarshalling message", err)
			}
			log.Println("messageBody>>>>>>", messageBody.UserEmail)

			services.SendMail("Loan Update", messageBody.Message, messageBody.UserEmail)
		}
	}()
	log.Println(" [*]Waiting for messages.. ")

	<-forever
}

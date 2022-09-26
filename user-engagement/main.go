package main

import (
	"fmt"
	"log"
	"user-engagement/rabbitmq"
)

func main() {

	rabbitMqConn, err := rabbitmq.Connect()
	if err != nil {
		log.Println("err", err)
	}

	ch, err := rabbitMqConn.Channel()

	if err != nil {
		log.Println("err at getting rabbitmq channel", err)
	}

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println("err at consume", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Rcvd msg: %s\n", d.Body)
		}
	}()

	log.Println("Rabbit: [*]Waiting for msgs[*]")
	<-forever
}

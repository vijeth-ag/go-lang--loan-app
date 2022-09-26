package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func Connect() (*amqp.Connection, error) {
	log.Println("connectng to rabbitMq...")

	RabbitMqConnection, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Println("err connecting to rabbitMq")
		return nil, err
	}

	return RabbitMqConnection, nil

}

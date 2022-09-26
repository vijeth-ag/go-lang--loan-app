package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

var RabbitmqConnection *amqp.Connection

func Connect() (*amqp.Connection, error) {
	log.Println("connecting to rabbit...")
	RabbitmqConnection, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Println("error connecting to rabbitmq", err)
		return nil, err
	}

	_ = RabbitmqConnection
	// defer RabbitmqConnection.Close()
	log.Println("Connected to rabbitmq")
	return RabbitmqConnection, nil
}

package clients

import (
	"encoding/json"
	"fmt"
	"loan/models"
	"log"
	"time"

	"github.com/streadway/amqp"
)

var RabbitMqConn *amqp.Connection

func init() {
	Connect()
}

func Connect() error {
	log.Println("connecting to rabbit...")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	RabbitMqConn = conn
	if err != nil {
		log.Println("error connecting to rabbitmq", err)
		return err
	}
	log.Println("Connected to rabbitmq")
	return nil
}

func PublishUserCommunicationMessage(user string, msg string) bool {
	message := models.UserCommunication{
		UserEmail: user,
		Message:   msg,
		DateTime:  fmt.Sprint(time.Now().UnixNano()),
	}

	ch, err := RabbitMqConn.Channel()
	if err != nil {
		log.Println("err", err)
		return false
	}
	q, err := ch.QueueDeclare(
		"UserCommunications",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("err decalring Q")
		return false
	}

	log.Println("q", q)

	messageTempData, err := json.Marshal(message)
	if err != nil {
		log.Println("err marshalling ", err)
		return false
	}

	err = ch.Publish(
		"",
		"UserCommunications",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        messageTempData,
		},
	)

	if err != nil {
		log.Println("err publishing user created msg")
		return false
	}

	return true
}

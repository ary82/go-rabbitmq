package main

import (
	"context"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	url := os.Getenv("AMQP_URL")
	conn, err := amqp091.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	q, err := channel.QueueDeclare(
		"testing",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.PublishWithContext(
		context.Background(),
		"",
		"testing",
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Text message"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("queue:", q)
	log.Println("Publish Success")
}

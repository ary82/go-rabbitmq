package main

import (
	"fmt"
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

	msgs, err := channel.Consume(
		"testing",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Println("Received:", string(msg.Body))
		}
	}()
	<-forever
}

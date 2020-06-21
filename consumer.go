package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer Application")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue", // queue string
		"",          // consumer string
		true,        // autoAck bool
		false,       // exclusive bool
		false,       // noLocal bool
		false,       // noWait bool
		nil,         // args amqp.Table
	) // (<-chan amqp.Delivery, error)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received Message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully connected to our RabbitMQ instance")
	fmt.Println(" [*] - waiting for messages")
	<-forever // block main func from exiting
}

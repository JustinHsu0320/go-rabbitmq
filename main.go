package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("GO RabbitMQ Start")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully Connected to out RabbitMQ Instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue", // name string
		false,       // durable bool
		false,       // autoDelete bool
		false,       // exclusive bool
		false,       // noWait bool
		nil,         // args amqp.Table
	) // (amqp.Queue, error)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(q)

	err = ch.Publish(
		"",          // exchange string
		"TestQueue", // key string
		false,       // mandatory bool
		false,       // immediate bool
		amqp.Publishing{ // msg amqp.Publishing
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	) // (error)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Successfully Publishing Message to Queue")

}

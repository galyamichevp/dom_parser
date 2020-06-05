package sockets

import (
	"fmt"
	"go-dom-parser/configs"
	"log"
	"os"

	"github.com/streadway/amqp"
)

//SetupRMQ - setup RMQ instance
func SetupRMQ(cfg *configs.Configuration) {
	rmqURL := configs.RMQURL(configs.BuildRMQConfig(cfg))

	conn, err := amqp.Dial(rmqURL)
	handleError(err, "Can't connect to AMQP")

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	amqpChannel.ExchangeDeclare(
		cfg.RMQ.Exchange,     // name of the exchange
		cfg.RMQ.ExchangeType, // type
		false,                // durable
		false,                // delete when complete
		false,                // internal
		false,                // noWait
		nil,                  // arguments
	)

	// create the queue if it doesn't already exist
	_, err = amqpChannel.QueueDeclare(cfg.RMQ.Queue, true, false, false, false, nil)
	handleError(err, "Could not declare `cfg.RMQ.Queue` queue")

	// err = amqpChannel.QueueBind(cfg.RMQ.Queue, "#", cfg.RMQ.Exchange, false, nil)
	err = amqpChannel.QueueBind(
		cfg.RMQ.Queue,      // name of the queue
		cfg.RMQ.RoutingKey, // bindingKey
		cfg.RMQ.Exchange,   // sourceExchange
		false,              // noWait
		nil,                // arguments
	)
	handleError(err, "Could not bind to `cfg.RMQ.Queue` queue")

	// prefetch 4x as many messages as we can handle at once
	prefetchCount := cfg.RMQ.Concurrency * 4
	err = amqpChannel.Qos(prefetchCount, 0, false)
	handleError(err, "Could not configure QoS")

	messageChannel, err := amqpChannel.Consume(
		cfg.RMQ.Queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Could not register consumer")

	// create a goroutine for the number of concurrent threads requested
	for i := 0; i < cfg.RMQ.Concurrency; i++ {
		fmt.Printf("Processing messages on thread %v...\n", i)
		go func() {
			for msg := range messageChannel {
				if handler(msg) {
					msg.Ack(false)
				} else {
					msg.Nack(false, true)
				}
			}
			fmt.Println("Rabbit consumer closed - critical Error")
			os.Exit(1)
		}()
	}
}

func handler(msg amqp.Delivery) bool {
	if msg.Body == nil {
		fmt.Println("Error, no message body!")
		return false
	}
	fmt.Println(string(msg.Body))
	return true
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

package sockets

import (
	"fmt"
	"go-dom-parser/configs"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//SetupRMQ - setup RMQ instance
func SetupRMQ(cfg *configs.Configuration) {
	xx := configs.RMQURL(configs.BuildRMQConfig(cfg))

	conn, err := amqp.Dial(xx)
	handleError(err, "Can't connect to AMQP")
	// defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	// defer amqpChannel.Close()

	// create the queue if it doesn't already exist
	queue, err := amqpChannel.QueueDeclare(cfg.RMQ.Queue, true, false, false, false, nil)
	handleError(err, "Could not declare `cfg.RMQ.Queue` queue")

	// prefetch 4x as many messages as we can handle at once
	prefetchCount := cfg.RMQ.Concurrency * 4
	err = amqpChannel.Qos(prefetchCount, 0, false)
	handleError(err, "Could not configure QoS")

	messageChannel, err := amqpChannel.Consume(
		queue.Name,
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
			for d := range messageChannel {
				log.Printf("Received a message: %s", d.Body)

				// processing here

				if err := d.Ack(false); err != nil {
					log.Printf("Error acknowledging message : %s", err)
				} else {
					log.Printf("Acknowledged message")
				}
			}
			fmt.Println("Rabbit consumer closed - critical Error")
			os.Exit(1)
		}()
	}
}

package sockets

import (
	"fmt"
	"go-dom-parser/configs"
	"log"
	"os"

	"github.com/streadway/amqp"
)

// Conn -
type Conn struct {
	Channel *amqp.Channel
}

//SetupRMQ - setup RMQ instance
func SetupRMQ(cfg *configs.Configuration) *Conn {
	rmqURL := configs.RMQURL(configs.BuildRMQConfig(cfg))

	conn, err := amqp.Dial(rmqURL)
	handleError(err, "Can't connect to AMQP")

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	amqpChannel.ExchangeDeclare(
		cfg.RMQ.ExchangeOut,     // name of the exchange
		cfg.RMQ.ExchangeTypeOut, // type
		false,                   // durable
		false,                   // delete when complete
		false,                   // internal
		false,                   // noWait
		nil,                     // arguments
	)

	// create the queue if it doesn't already exist
	_, err = amqpChannel.QueueDeclare(cfg.RMQ.QueueOut, true, false, false, false, nil)
	handleError(err, "Could not declare `cfg.RMQ.Queue` queue")

	// err = amqpChannel.QueueBind(cfg.RMQ.Queue, "#", cfg.RMQ.Exchange, false, nil)
	err = amqpChannel.QueueBind(
		cfg.RMQ.QueueOut,      // name of the queue
		cfg.RMQ.RoutingKeyOut, // bindingKey
		cfg.RMQ.ExchangeOut,   // sourceExchange
		false,                 // noWait
		nil,                   // arguments
	)
	handleError(err, "Could not bind to `cfg.RMQ.Queue` queue")

	return &Conn{
		Channel: amqpChannel,
	}
}

// Publish - publish message to RMQ
func (conn *Conn) Publish(cfg *configs.Configuration, body []byte) error {
	log.Printf("send message to exchange")

	message := amqp.Publishing{
		//DeliveryMode: amqp.Persistent,
		//Timestamp:    time.Now(),
		//ContentType:  "text/plain",
		//Body:         []byte(body),
		Headers:         amqp.Table{},
		ContentType:     "text/plain",
		ContentEncoding: "",
		Body:            body,
		DeliveryMode:    1, // 1=non-persistent, 2=persistent
		Priority:        0, // 0-9
	}

	return conn.Channel.Publish(
		cfg.RMQ.ExchangeOut,   // publish to an exchange
		cfg.RMQ.RoutingKeyOut, // routing to 0 or more queues
		false,                 // mandatory
		false,                 // immediate
		message,
	)
}

// Subscribe - subscribe to RMQ
func (conn *Conn) Subscribe(cfg *configs.Configuration) error {
	messageChannel, err := conn.Channel.Consume(
		cfg.RMQ.QueueIn,
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

	return nil
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

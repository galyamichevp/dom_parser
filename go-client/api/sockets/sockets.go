package sockets

import (
	"go-dom-parser/configs"
	"log"

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
		cfg.RMQ.Exchange,   // publish to an exchange
		cfg.RMQ.RoutingKey, // routing to 0 or more queues
		false,              // mandatory
		false,              // immediate
		message,
	)

}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

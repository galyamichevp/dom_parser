package sockets

import (
	"fmt"
	"go-dom-parser/configs"
	"log"
	"os"

	"github.com/streadway/amqp"
)

//SetupRMQ - setup RMQ instance
func SetupRMQ(cfg *configs.Configuration) *Conn {
	rmqURL := configs.RMQURL(configs.BuildRMQConfig(cfg))

	conn, err := amqp.Dial(rmqURL)
	handleError(err, "Can't connect to AMQP")

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	// income queue with web pages
	buildChannel(amqpChannel, cfg.RMQ.ExchangeIn, cfg.RMQ.ExchangeTypeIn, cfg.RMQ.QueueIn, cfg.RMQ.RoutingKeyIn, cfg.RMQ.Concurrency)
	// outcome queue to send parse result
	//buildChannel(amqpChannel, cfg.RMQ.ExchangeOut, cfg.RMQ.ExchangeTypeOut, cfg.RMQ.QueueOut, cfg.RMQ.RoutingKeyOut, cfg.RMQ.Concurrency)

	return &Conn{
		Channel:    amqpChannel,
		Processors: make(map[string][]chan string),
	}
}

// Publish - publish message to RMQ
func (conn *Conn) Publish(cfg *configs.Configuration, body []byte) error {

	e := "ctrl"

	if _, ok := conn.Processors[e]; ok {
		for _, handler := range conn.Processors[e] {
			go func(handler chan string, body []byte, cfg *configs.Configuration) {
				for {
					select {
					case body := <-handler:
						//fmt.Println("processed res: " + body)
						fmt.Println("Sent request to: " + cfg.RMQ.ExchangeOut)

						message := amqp.Publishing{
							//DeliveryMode: amqp.Persistent,
							//Timestamp:    time.Now(),
							//ContentType:  "text/plain",
							//Body:         []byte(body),
							Headers:         amqp.Table{},
							ContentType:     "text/plain",
							ContentEncoding: "",
							Body:            []byte(body),
							DeliveryMode:    1, // 1=non-persistent, 2=persistent
							Priority:        0, // 0-9
						}

						conn.Channel.Publish(
							cfg.RMQ.ExchangeOut,   // publish to an exchange
							cfg.RMQ.RoutingKeyOut, // routing to 0 or more queues
							false,                 // mandatory
							false,                 // immediate
							message,
						)
						// return
						// case <-time.After(10 * time.Second):
						// 	fmt.Println("FAIL res: ")
						// 	return
					}
				}
			}(handler, body, cfg)
		}
	}

	return nil
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

				fmt.Println("INCOME msg ... ")

				if conn.handler(msg) {
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

func buildChannel(amqpChannel *amqp.Channel, exchange, exchangeType, queue, routingKey string, concurrency int) {
	amqpChannel.ExchangeDeclare(
		exchange,     // name of the exchange
		exchangeType, // type
		false,        // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)

	// create the queue if it doesn't already exist
	_, err := amqpChannel.QueueDeclare(queue, true, false, false, false, nil)
	handleError(err, "Could not declare `cfg.RMQ.Queue` queue")

	// err = amqpChannel.QueueBind(cfg.RMQ.Queue, "#", cfg.RMQ.Exchange, false, nil)
	err = amqpChannel.QueueBind(
		queue,      // name of the queue
		routingKey, // bindingKey
		exchange,   // sourceExchange
		false,      // noWait
		nil,        // arguments
	)
	handleError(err, "Could not bind to `cfg.RMQ.Queue` queue")

	// prefetch 4x as many messages as we can handle at once
	prefetchCount := concurrency * 4
	err = amqpChannel.Qos(prefetchCount, 0, false)
	handleError(err, "Could not configure QoS")
}

//handler handle queue message
func (conn *Conn) handler(msg amqp.Delivery) bool {
	if msg.Body == nil {
		fmt.Println("Error, no message body!")
		return false
	}

	conn.emit("test", string(msg.Body))

	return true
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

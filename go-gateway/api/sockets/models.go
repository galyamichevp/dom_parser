package sockets

import "github.com/streadway/amqp"

type OutPayload struct {
	Content string
	Type    string
	Marker  string
}

type InPayload struct {
	Content string
	Type    string
	Marker  string
}

// Conn -
type Conn struct {
	Channel    *amqp.Channel
	Processors map[string][]chan string
}

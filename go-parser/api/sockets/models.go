package sockets

import (
	"go-dom-parser/configs"

	"github.com/streadway/amqp"
)

type InPayload struct {
	Content string
	Type    string
	Marker  string
}

type OutPayload struct {
	Content string
	Type    string
	Marker  string
}

// Conn -
type Conn struct {
	Channel    *amqp.Channel
	Processors map[string][]chan string
	Cfg        *configs.Configuration
}

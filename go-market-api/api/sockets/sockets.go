package sockets

import (
	"go-dom-parser/configs"
)

//SetupSocket - setup socket instance
func SetupSocket(cfg *configs.Configuration) *Socket {
	return &Socket{
		Channels: make(map[string][]chan string),
	}
}

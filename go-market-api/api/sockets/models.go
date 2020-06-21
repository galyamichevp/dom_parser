package sockets

// Socket - allows to build up interconnection with processors and api
type Socket struct {
	Channels map[string][]chan string
}

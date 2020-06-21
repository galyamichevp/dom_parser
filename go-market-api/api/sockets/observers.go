package sockets

// AddProcessor adds an event listener
func (socket *Socket) AddChannel(e string, ch chan string) {
	if _, ok := socket.Channels[e]; ok {
		socket.Channels[e] = append(socket.Channels[e], ch)
	} else {
		socket.Channels[e] = []chan string{ch}
	}
}

// RemoveProcessor removes an event listener
func (socket *Socket) RemoveChannel(e string, ch chan string) {
	if _, ok := socket.Channels[e]; ok {
		for i := range socket.Channels[e] {
			if socket.Channels[e][i] == ch {
				socket.Channels[e] = append(socket.Channels[e][:i], socket.Channels[e][i+1:]...)
				break
			}
		}
	}
}

// Emit emits an event on the Dog struct instance
func (socket *Socket) emit(e string, response string) {
	if _, ok := socket.Channels[e]; ok {
		for _, handler := range socket.Channels[e] {
			go func(handler chan string, response string) {
				handler <- response
			}(handler, response)
		}
	}
}

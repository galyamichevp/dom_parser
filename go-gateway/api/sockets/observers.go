package sockets

// AddProcessor adds an event listener
func (conn *Conn) AddProcessor(e string, ch chan string) {
	if conn.Processors == nil {
		conn.Processors = make(map[string][]chan string)
	}
	if _, ok := conn.Processors[e]; ok {
		conn.Processors[e] = append(conn.Processors[e], ch)
	} else {
		conn.Processors[e] = []chan string{ch}
	}
}

// RemoveProcessor removes an event listener
func (conn *Conn) RemoveProcessor(e string, ch chan string) {
	if _, ok := conn.Processors[e]; ok {
		for i := range conn.Processors[e] {
			if conn.Processors[e][i] == ch {
				conn.Processors[e] = append(conn.Processors[e][:i], conn.Processors[e][i+1:]...)
				break
			}
		}
	}
}

// Emit emits an event on the Dog struct instance
func (conn *Conn) emit(e string, response string) {
	if _, ok := conn.Processors[e]; ok {
		for _, handler := range conn.Processors[e] {
			go func(handler chan string, response string) {
				handler <- response
			}(handler, response)
		}
	}
}

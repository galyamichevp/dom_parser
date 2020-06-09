package core

import (
	"fmt"
)

type Processor struct {
	ProcessorChan chan string
}

func New() *Processor {
	out := make(chan string)

	return &Processor{
		ProcessorChan: out,
	}
}

// Run - ...
func (p *Processor) Run() {
	go func() {
		for {
			msg := <-p.ProcessorChan
			fmt.Println("INFO: " + msg)
		}
	}()
}

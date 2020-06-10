package core

import (
	"encoding/json"
	"fmt"
	"log"
)

type Processor struct {
	ProcessorChan chan string
	SPBStocks     []SPBStock
}

func New() *Processor {
	out := make(chan string)

	return &Processor{
		ProcessorChan: out,
	}
}

type PayloadRMQ struct {
	Content string
	Type    string
}

type SPBStock struct {
	Id       string
	Marker   string
	Title    string
	Code1    string
	Code2    string
	Count    string
	Price    string
	Currency string
	Date     string
	Note     string
}

// Run - ...
func (p *Processor) Run() {
	go func() {
		for {
			msg := <-p.ProcessorChan
			fmt.Println("INFO: " + msg)

			payload := &PayloadRMQ{}
			err := json.Unmarshal([]byte(msg), payload)
			if err != nil {
				log.Printf("ERROR: fail unmarshl: %s", err.Error)
			}

			if payload.Type == "result.spb" {
				p.loadSpbStocks(payload.Content)
			}

			fmt.Println("INFO: result content ... " + payload.Content)
		}
	}()
}

func (p *Processor) loadSpbStocks(data string) {
	var arr []SPBStock
	_ = json.Unmarshal([]byte(data), &arr)

	p.SPBStocks = append(p.SPBStocks, arr...)
}

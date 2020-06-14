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
	Marker  string
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

	Today   []string
	Days30  []string
	Days90  []string
	Days180 []string

	Percent float64
}

// Run - ...
func (p *Processor) Run() {
	go func() {
		for {
			msg := <-p.ProcessorChan
			// fmt.Println("INFO: " + msg)

			payload := &PayloadRMQ{}
			err := json.Unmarshal([]byte(msg), payload)
			if err != nil {
				log.Printf("ERROR: fail unmarshl: %s", err.Error)
			}

			if payload.Type == "result.spb.foreignexchange" {
				p.LoadSpbStocks(payload.Content)
			}

			if payload.Type == "result.marketbeat.ratings" {
				p.loadMarketBeatStocks(payload.Content)
				fmt.Println("INFO: result MarketBeat for marker ... " + payload.Marker)
			}

		}
	}()
}

func (p *Processor) LoadSpbStocks(data string) {
	fmt.Println("INFO: load spb stocks ...")

	var arr []SPBStock
	_ = json.Unmarshal([]byte(data), &arr)

	// fmt.Println("INFO: load spb stocks ..." + strconv.Itoa(len(arr)))

	p.SPBStocks = append(p.SPBStocks, arr...)
}

type MarketBeatStock struct {
	Marker  string
	Today   []string
	Days30  []string
	Days90  []string
	Days180 []string
}

func (p *Processor) loadMarketBeatStocks(data string) {
	var mb MarketBeatStock
	_ = json.Unmarshal([]byte(data), &mb)

	for i, val := range p.SPBStocks {
		if val.Marker == mb.Marker && mb.Today != nil {
			p.SPBStocks[i].Today = mb.Today
			p.SPBStocks[i].Days30 = mb.Days30
			p.SPBStocks[i].Days90 = mb.Days90
			p.SPBStocks[i].Days180 = mb.Days180
		}
	}
}

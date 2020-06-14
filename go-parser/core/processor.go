package core

import (
	"encoding/json"
	"fmt"
	"go-dom-parser/api/sockets"
	"log"
)

func New() *Processor {
	out := make(chan string)
	cout := make(chan string)

	return &Processor{
		ProcessorChan:    out,
		ProcessorChanOut: cout,
	}
}

// Run - ...
func (p *Processor) Run() {
	go func() {
		for {
			msg := <-p.ProcessorChan
			fmt.Println("INFO: income message")

			payload := &sockets.InPayload{}
			err := json.Unmarshal([]byte(msg), payload)
			if err != nil {
				log.Printf("ERROR: fail unmarshl: %s", err.Error)
			}

			var result = &sockets.OutPayload{}
			if payload.Type == "resource.spbexchange.foreignshares" {
				result = spbParseForeignShare(payload)
			}

			if payload.Type == "resource.marketbeat.ratings" {
				result = marketBeatParseRating(payload)
			}

			out, err := json.Marshal(result)
			if err != nil {
				log.Printf("ERROR: fail marshl: %s", err.Error)
			}

			p.ProcessorChanOut <- string(out)
		}
	}()
}

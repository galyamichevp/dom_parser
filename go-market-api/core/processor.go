package core

import (
	"go-dom-parser/api/sockets"
	"go-dom-parser/domain"
)

//SetupProcessor - initialize processor
func SetupProcessor(storage *domain.Storage, socket *sockets.Socket) *Processor {

	return &Processor{
		Socket:  socket,
		Storage: storage,
	}
}

// RunProcessor - run processor to listenin incomin payload messages
func (processor *Processor) RunProcessor() {

	for key, channel := range processor.Socket.Channels {
		go func(key string, channel []chan string) {
			if len(channel) <= 0 {
				return
			}

			for {
				select {
				case payload := <-channel[0]:

					if key == "spbexchange.loadsymbols" {
						// fmt.Println("INFO: symbol loading ...")

						processor.spbExchangeParseSymbol(payload)
					}

					if key == "marketbeat.loadanalytics" {
						// fmt.Println("INFO: symbol loading ...")

						processor.marketBeatParseAnalytics(payload)
					}

					if key == "nasdaq.loadinfo" {
						// fmt.Println("INFO: symbol loading ...")

						processor.nasdaqParseInfo(payload)
					}

					if key == "nasdaq.loadsummary" {
						// fmt.Println("INFO: symbol loading ...")

						processor.nasdaqParseSummary(payload)
					}

					if key == "nasdaq.loadrealtime" {
						// fmt.Println("INFO: symbol loading ...")

						processor.nasdaqParseRealTime(payload)
					}

					if key == "nasdaq.loadhistory" {
						// fmt.Println("INFO: symbol loading ...")

						processor.nasdaqParseHistory(payload)
					}

				}
			}
		}(key, channel)
	}
}

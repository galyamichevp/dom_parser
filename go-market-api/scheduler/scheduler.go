package scheduler

import (
	"go-dom-parser/api/sockets"
	"go-dom-parser/domain"
	"go-dom-parser/scheduler/jobs"
	"time"
)

//SetupScheduler - initialize scheduler
func SetupScheduler(storage *domain.Storage, socket *sockets.Socket) *Scheduler {

	// ... spb exchange load symbols job
	symbolChan := make(chan string)
	socket.AddChannel("spbexchange.loadsymbols", symbolChan)
	// ...

	// ... marketbeat load analytics job
	analyticChan := make(chan string)
	socket.AddChannel("marketbeat.loadanalytics", analyticChan)
	// ...

	// ... nasdaq load info job
	infoChan := make(chan string)
	socket.AddChannel("nasdaq.loadinfo", infoChan)
	// ...

	// ... nasdaq load summary job
	summaryChan := make(chan string)
	socket.AddChannel("nasdaq.loadsummary", summaryChan)
	// ...

	// ... nasdaq load realtime job
	realTimeChan := make(chan string)
	socket.AddChannel("nasdaq.loadrealtime", realTimeChan)
	// ...

	// ... nasdaq load history job
	historyChan := make(chan string)
	socket.AddChannel("nasdaq.loadhistory", historyChan)
	// ...

	jobs := []jobs.Job{
		jobs.SpbExchange{
			SymbolChan: symbolChan,
		},
		jobs.MarketBeat{
			AnalyticsChan: analyticChan,
			Storage:       storage,
		},
		jobs.Nasdaq{
			InfoChan:     infoChan,
			SummaryChan:  summaryChan,
			RealTimeChan: realTimeChan,
			HistoryChan:  historyChan,
			Storage:      storage,
		},
	}

	return &Scheduler{
		Jobs:    jobs,
		Storage: storage,
	}
}

//RunJobs - run spb exchange job
func (scheduler *Scheduler) RunJobs() {
	for _, job := range scheduler.Jobs {
		time.Sleep(5 * time.Second)
		job.RunJob()
	}
}

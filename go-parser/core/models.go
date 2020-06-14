package core

type Processor struct {
	ProcessorChan    chan string
	ProcessorChanOut chan string
}

type MarketBeatStock struct {
	Marker  string
	Today   []string
	Days30  []string
	Days90  []string
	Days180 []string
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

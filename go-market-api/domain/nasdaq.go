package domain

type NasdaqInfo struct {
	Data    NasdaqInfoData `json:"data"`
	Message interface{}    `json:"message"`
	Status  Status         `json:"status"`
}

type NasdaqInfoData struct {
	Symbol         string      `json:"symbol"`
	CompanyName    string      `json:"companyName"`
	StockType      string      `json:"stockType"`
	Exchange       string      `json:"exchange"`
	IsNasdaqListed bool        `json:"isNasdaqListed"`
	IsNasdaq100    bool        `json:"isNasdaq100"`
	IsHeld         bool        `json:"isHeld"`
	PrimaryData    PrimaryData `json:"primaryData"`
	SecondaryData  interface{} `json:"secondaryData"`
	KeyStats       KeyStats    `json:"keyStats"`
	MarketStatus   string      `json:"marketStatus"`
	AssetClass     string      `json:"assetClass"`
}

type KeyStats struct {
	Volume        MarketCap `json:"Volume"`
	PreviousClose MarketCap `json:"PreviousClose"`
	OpenPrice     MarketCap `json:"OpenPrice"`
	MarketCap     MarketCap `json:"MarketCap"`
}

type MarketCap struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type PrimaryData struct {
	LastSalePrice      string `json:"lastSalePrice"`
	NetChange          string `json:"netChange"`
	PercentageChange   string `json:"percentageChange"`
	DeltaIndicator     string `json:"deltaIndicator"`
	LastTradeTimestamp string `json:"lastTradeTimestamp"`
	IsRealTime         bool   `json:"isRealTime"`
}

type Status struct {
	RCode            int64       `json:"rCode"`
	BCodeMessage     interface{} `json:"bCodeMessage"`
	DeveloperMessage interface{} `json:"developerMessage"`
}

type NasdaqSummary struct {
	Data    NasdaqSummaryData `json:"data"`
	Message interface{}       `json:"message"`
	Status  Status            `json:"status"`
}

type NasdaqSummaryData struct {
	Symbol         string      `json:"symbol"`
	SummaryData    SummaryData `json:"summaryData"`
	AssetClass     string      `json:"assetClass"`
	AdditionalData interface{} `json:"additionalData"`
}

type SummaryData struct {
	Exchange           AnnualizedDividend `json:"Exchange"`
	Sector             AnnualizedDividend `json:"Sector"`
	Industry           AnnualizedDividend `json:"Industry"`
	OneYrTarget        AnnualizedDividend `json:"OneYrTarget"`
	TodayHighLow       AnnualizedDividend `json:"TodayHighLow"`
	ShareVolume        AnnualizedDividend `json:"ShareVolume"`
	AverageVolume      AnnualizedDividend `json:"AverageVolume"`
	PreviousClose      AnnualizedDividend `json:"PreviousClose"`
	FiftTwoWeekHighLow AnnualizedDividend `json:"FiftTwoWeekHighLow"`
	MarketCap          AnnualizedDividend `json:"MarketCap"`
	// PERatio             BetaX              `json:"PERatio"`
	ForwardPE1Yr        AnnualizedDividend `json:"ForwardPE1Yr"`
	EarningsPerShare    AnnualizedDividend `json:"EarningsPerShare"`
	AnnualizedDividend  AnnualizedDividend `json:"AnnualizedDividend"`
	ExDividendDate      AnnualizedDividend `json:"ExDividendDate"`
	DividendPaymentDate AnnualizedDividend `json:"DividendPaymentDate"`
	Yield               AnnualizedDividend `json:"Yield"`
	// Beta                BetaX              `json:"Beta"`
}

type AnnualizedDividend struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type BetaX struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}

type NasdaqRealTime struct {
	Data    RealTimeData `json:"data"`
	Message interface{}  `json:"message"`
	Status  Status       `json:"status"`
}

type RealTimeData struct {
	Symbol       string            `json:"symbol"`
	TotalRecords int64             `json:"totalRecords"`
	Offset       int64             `json:"offset"`
	Limit        int64             `json:"limit"`
	Headers      RealTimeHeaders   `json:"headers"`
	Rows         []RealTimeHeaders `json:"rows"`
}

type RealTimeHeaders struct {
	NLSTime        string `json:"nlsTime"`
	NLSPrice       string `json:"nlsPrice"`
	NLSShareVolume string `json:"nlsShareVolume"`
}

// ... history
type NasdaqHistory struct {
	Data    NasdaqHistoryData `json:"data"`
	Message interface{}       `json:"message"`
	Status  Status            `json:"status"`
}

type NasdaqHistoryData struct {
	Symbol           string               `json:"symbol"`
	Company          string               `json:"company"`
	TimeAsOf         string               `json:"timeAsOf"`
	IsNasdaq100      bool                 `json:"isNasdaq100"`
	LastSalePrice    string               `json:"lastSalePrice"`
	NetChange        string               `json:"netChange"`
	PercentageChange string               `json:"percentageChange"`
	DeltaIndicator   string               `json:"deltaIndicator"`
	PreviousClose    string               `json:"previousClose"`
	Chart            []NasdaqHistoryChart `json:"chart"`
	Events           interface{}          `json:"events"`
}

type NasdaqHistoryChart struct {
	Z NasdaqHistoryZ `json:"z"`
	X int64          `json:"x"`
	Y float64        `json:"y"`
}

type NasdaqHistoryZ struct {
	High     string `json:"high"`
	Low      string `json:"low"`
	Open     string `json:"open"`
	Close    string `json:"close"`
	Volume   string `json:"volume"`
	DateTime string `json:"dateTime"`
	Value    string `json:"value"`
}

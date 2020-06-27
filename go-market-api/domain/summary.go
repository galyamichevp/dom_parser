package domain

// Summary - summary entity
type Summary struct {
	ID                 string  `json:"id"`
	Title              string  `json:"title"`
	Symbol             string  `json:"symbol"`
	Sector             string  `json:"sector"`
	Industry           string  `json:"industry"`
	TodayHigh          float64 `json:"todayHigh"`
	TodayLow           float64 `json:"todayLow"`
	TodayVolatility    float64 `json:"todayVolatility"`
	ShareVolume        string  `json:"shareVolume"`
	AverageVolume      string  `json:"averageVolume"`
	PreviousClose      string  `json:"previousClose"`
	FiftTwoWeekHighLow string  `json:"fiftTwoWeekHighLow"`
	EarningsPerShare   float64 `json:"earningsPerShare"`
}

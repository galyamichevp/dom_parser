package domain

// Summary - summary entity
type Summary struct {
	ID                 string  `json:"id"`
	Title              string  `json:"title"`
	Symbol             string  `json:"symbol"`
	Sector             string  `json:"sector"`
	Industry           string  `json:"industry"`
	TodayHighLow       string  `json:"todayHighLow"`
	ShareVolume        string  `json:"shareVolume"`
	AverageVolume      string  `json:"averageVolume"`
	PreviousClose      string  `json:"previousClose"`
	FiftTwoWeekHighLow string  `json:"fiftTwoWeekHighLow"`
	EarningsPerShare   float64 `json:"earningsPerShare"`
}

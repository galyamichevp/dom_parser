package domain

// Summary - summary entity
type Summary struct {
	ID               string  `json:"id"`
	Title            string  `json:"title"`
	Symbol           string  `json:"symbol"`
	Sector           string  `json:"sector"`
	Industry         string  `json:"industry"`
	TodayHigh        float64 `json:"todayHigh"`
	TodayLow         float64 `json:"todayLow"`
	ShareVolume      string  `json:"shareVolume"`
	AverageVolume    string  `json:"averageVolume"`
	PreviousClose    float64 `json:"previousClose"`
	FiftTwoWeekHigh  float64 `json:"fiftTwoWeekHigh"`
	FiftTwoWeekLow   float64 `json:"fiftTwoWeekLow"`
	EarningsPerShare float64 `json:"earningsPerShare"`
	PERatio          float64 `json:"peRatio"`
}

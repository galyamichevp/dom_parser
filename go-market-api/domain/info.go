package domain

// Info - summary entity
type Info struct {
	ID               string  `json:"id"`
	Title            string  `json:"title"`
	Symbol           string  `json:"symbol"`
	CompanyName      string  `json:"companyName"`
	StockType        string  `json:"stockType"`
	Exchange         string  `json:"exchange"`
	Volume           string  `json:"volume"`
	PreviousClose    string  `json:"previousClose"`
	OpenPrice        string  `json:"openPrice"`
	LastSalePrice    float64 `json:"lastSalePrice"`
	NetChange        string  `json:"netChange"`
	PercentageChange float64 `json:"percentageChange"`
	DeltaIndicator   string  `json:"deltaIndicator"`
}

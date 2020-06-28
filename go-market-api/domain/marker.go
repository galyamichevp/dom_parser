package domain

// Marker - ...
type Marker struct {
	ID     string  `json:"id"`
	Symbol string  `json:"symbol"`
	Title  string  `json:"title"`
	FValue float64 `json:"fValue"`
	SValue string  `json:"sValue"`
	BValue bool    `json:"bValue"`
}

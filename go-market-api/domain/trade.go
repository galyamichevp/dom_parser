package domain

import "time"

// Trade - trade entity
type Trade struct {
	ID     uint      `json:"id"`
	Symbol string    `json:"symbol"`
	Price  float64   `json:"price"`
	Volume float64   `json:"volume"`
	Time   time.Time `json:"time"`
}

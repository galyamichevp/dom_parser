package domain

import "time"

// Chart - chart entity
type Chart struct {
	High     float64   `json:"high"`
	Low      float64   `json:"low"`
	Open     float64   `json:"open"`
	Close    float64   `json:"close"`
	Volume   float64   `json:"volume"`
	DateTime time.Time `json:"dateTime"`
	Value    float64   `json:"value"`
}

package domain

//Resource - structure to processing
type Resource struct {
	ID      uint   `json:"id"`
	Payload string `json:"payload"`
	Pattern string `json:"pattern"`
}

package domain

// History - historical entity
type History struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Symbol string  `json:"symbol"`
	Chart  []Chart `json:"chart"`
}

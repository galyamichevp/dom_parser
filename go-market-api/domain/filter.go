package domain

// Filter - filter entity
type Filter struct {
	IsActive        bool   `json:"isActive"`
	Symbol          string `json:"symbol"`
	LoTargetPercent string `json:"loTargetPercent"`
}

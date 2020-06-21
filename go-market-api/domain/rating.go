package domain

// Rating - rating entity
type Rating struct {
	ID                   string  `json:"id"`
	Title                string  `json:"title"`
	Symbol               string  `json:"symbol"`
	TragetPrice          string  `json:"targetPrice"`
	TragetPercent        float64 `json:"targetPercent"`
	ConsensusRating      string  `json:"consensusRating"`
	ConsensusRatingScore string  `json:"consensusRatingScore"`
	ConsensusBreakdown   string  `json:"consensusBreakdown"`
}

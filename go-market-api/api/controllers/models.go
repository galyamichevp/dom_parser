package controllers

import "go-dom-parser/domain"

// Controller - container to coomunicate with other infrastructure components
type Controller struct {
	Storage *domain.Storage
}

// ResourcesGetRequest - query model allowed to get requests
type ResourcesGetRequest struct {
	Curl string `form:"curl"`
	Type string `form:"type"`
	Verb string `form:"verb"`
}

// SymbolsGetResponse - ...
type SymbolsGetResponse struct {
	Symbols    []domain.Symbol `json:"symbols"`
	Filters    []string        `json:"filters"`
	TotalPages int             `json:"totalPages"`
}

// SymbolsGetRequest - query model allowed to get symbols
type SymbolsGetRequest struct {
	SortTargetPercent string    `form:"sortTargetPercent"`
	TargetPercents    []float64 `form:"targetPercents[]" binding:"required"`
	DeltaPercents     []float64 `form:"deltaPercents[]" binding:"required"`
	Page              int       `form:"page"`
	PageSize          int       `form:"pageSize"`
}

// FilterPostRequest - query model allowed to post filter
type FilterPostRequest struct {
	Symbol string `json:"symbol"`
	State  bool   `json:"state"`
}

// FilterDeleteRequest - query model allowed to delete filter
type FilterDeleteRequest struct {
	Symbol string `form:"symbol"`
}

// SyncPostRequest - query model allowed to post sync
type SyncPostRequest struct {
	Source string `json:"source"`
	State  bool   `json:"state"`
}

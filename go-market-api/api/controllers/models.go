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

// SymbolsGetRequest - query model allowed to get symbols
type SymbolsGetRequest struct {
	SortByPercent string `form:"sort_percent"`
	PercentLimit  string `form:"percent_limit"`
}

// FilterPostRequest - query model allowed to post filter
type FilterPostRequest struct {
	Symbols []string `json:"symbols"`
}

// FilterDeleteRequest - query model allowed to delete filter
type FilterDeleteRequest struct {
	Symbol string `form:"symbol"`
}

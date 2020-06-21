package controllers

import "go-dom-parser/domain"

// SortAscBySymbolId - comparator for Symbols
type SortAscBySymbolId []domain.Symbol

// Len - ByPercent comparator function
func (a SortAscBySymbolId) Len() int { return len(a) }

// Swap - ByPercent comparator function
func (a SortAscBySymbolId) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less - ByPercent comparator function
func (a SortAscBySymbolId) Less(i, j int) bool {
	return a[i].ID > a[j].ID
}

// SortAscByRatingPercent - comparator for SPBStock
type SortAscByRatingPercent []domain.Symbol

// Len - ByPercent comparator function
func (a SortAscByRatingPercent) Len() int { return len(a) }

// Swap - ByPercent comparator function
func (a SortAscByRatingPercent) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less - ByPercent comparator function
func (a SortAscByRatingPercent) Less(i, j int) bool {
	return a[i].Ratings != nil && a[j].Ratings != nil && a[i].Ratings["marketbeat"].TragetPercent < a[j].Ratings["marketbeat"].TragetPercent
}

// SortDescByRatingPercent - comparator for SPBStock
type SortDescByRatingPercent []domain.Symbol

// Len - ByPercent comparator function
func (a SortDescByRatingPercent) Len() int { return len(a) }

// Swap - ByPercent comparator function
func (a SortDescByRatingPercent) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less - ByPercent comparator function
func (a SortDescByRatingPercent) Less(i, j int) bool {
	if a[i].Ratings == nil {
		return false
	}

	if a[j].Ratings == nil {
		return true
	}

	return a[i].Ratings["marketbeat"].TragetPercent > a[j].Ratings["marketbeat"].TragetPercent
}

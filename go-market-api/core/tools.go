package core

import "go-dom-parser/domain"

// SortByDateTrade - comparator for Trade
type SortByDateTrade []domain.Trade

// Len - SortByDateTrade comparator function
func (a SortByDateTrade) Len() int { return len(a) }

// Swap - SortByDateTrade comparator function
func (a SortByDateTrade) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less - SortByDateTrade comparator function
func (a SortByDateTrade) Less(i, j int) bool {
	return a[i].Time.After(a[j].Time)
}

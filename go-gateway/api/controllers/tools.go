package controllers

import "go-dom-parser/core"

// ByPercent - comparator for SPBStock
type ByPercent []core.SPBStock

// Len - ByPercent comparator function
func (a ByPercent) Len() int { return len(a) }

// Swap - ByPercent comparator function
func (a ByPercent) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less - ByPercent comparator function
func (a ByPercent) Less(i, j int) bool { return a[i].Percent > a[j].Percent }

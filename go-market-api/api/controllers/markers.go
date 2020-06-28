package controllers

import "go-dom-parser/domain"

// BuildMarkers - build markers by pre-defined data
func (r *SymbolsGetResponse) BuildMarkers() {

	for i, _ := range r.Symbols {
		r.Symbols[i].Markers["x"] = domain.Marker{
			Symbol: "DDD",
		}
	}
}

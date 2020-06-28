package domain

// Symbol - symbol entity
type Symbol struct {
	ID          uint64               `json:"id"`
	Title       string               `json:"title"`
	Symbol      string               `json:"symbol"`
	Price       string               `json:"price"`
	Currency    string               `json:"currency"`
	Description string               `json:"description"`
	Ratings     map[string]Rating    `json:"ratings"`
	Infos       map[string]Info      `json:"infos"`
	Summaries   map[string]Summary   `json:"summaries"`
	Indicators  map[string]Indicator `json:"indicators"`
	Trades      map[string][]Trade   `json:"trades"`
	Histories   map[string]History   `json:"histories"`
	Markers     map[string]Marker    `json:"markers"`
}

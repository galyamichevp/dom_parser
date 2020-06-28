package domain

import "sync"

// Storage - global in-memory storage
type Storage struct {
	ID           uint              `json:"id"`
	Title        string            `json:"title"`
	Description  string            `json:"description"`
	Symbols      map[string]Symbol `json:"symbols"`
	Filters      map[string]Filter `json:"filters"`
	Locker       sync.RWMutex
	FilterLocker sync.RWMutex
	Sync         map[string]bool `json:"sync"`
}

func (s *Storage) Init() {
	s.Symbols = make(map[string]Symbol)
	s.Filters = make(map[string]Filter)
	s.Sync = make(map[string]bool)
}

func (s *Storage) GetSymbols() map[string]Symbol {
	s.Locker.RLock()
	defer s.Locker.RUnlock()

	return s.Symbols
}

func (s *Storage) GetSymbol(symbol string) Symbol {
	s.Locker.RLock()
	defer s.Locker.RUnlock()

	data := s.Symbols[symbol]

	if data.Ratings == nil {
		data.Ratings = make(map[string]Rating)
	}
	if data.Infos == nil {
		data.Infos = make(map[string]Info)
	}
	if data.Summaries == nil {
		data.Summaries = make(map[string]Summary)
	}
	if data.Indicators == nil {
		data.Indicators = make(map[string]Indicator)
	}
	if data.Trades == nil {
		data.Trades = make(map[string][]Trade)
	}
	if data.Histories == nil {
		data.Histories = make(map[string]History)
	}
	if data.Markers == nil {
		data.Markers = make(map[string]Marker)
	}

	return data
}

func (s *Storage) GetSymbolUnsafe(symbol string) Symbol {
	data := s.Symbols[symbol]

	if data.Ratings == nil {
		data.Ratings = make(map[string]Rating)
	}
	if data.Infos == nil {
		data.Infos = make(map[string]Info)
	}
	if data.Summaries == nil {
		data.Summaries = make(map[string]Summary)
	}
	if data.Indicators == nil {
		data.Indicators = make(map[string]Indicator)
	}
	if data.Trades == nil {
		data.Trades = make(map[string][]Trade)
	}
	if data.Histories == nil {
		data.Histories = make(map[string]History)
	}
	if data.Markers == nil {
		data.Markers = make(map[string]Marker)
	}

	return data
}

func (s *Storage) GetSymbolsKeys() []string {
	s.Locker.RLock()
	defer s.Locker.RUnlock()

	var keys []string
	for key, _ := range s.Symbols {
		keys = append(keys, key)
	}

	return keys
}

func (s *Storage) SetSymbol(symbol string, data Symbol) {
	s.Locker.Lock()
	defer s.Locker.Unlock()

	s.Symbols[symbol] = data
}

func (s *Storage) SetSymbolUnsafe(symbol string, data Symbol) {
	s.Symbols[symbol] = data
}

func (s *Storage) GetFilterSymbol(symbol string) Filter {
	s.Locker.RLock()
	defer s.Locker.RUnlock()

	return s.Filters[symbol]
}

func (s *Storage) ResetFilters() {
	s.FilterLocker.Lock()
	defer s.FilterLocker.Unlock()

	s.Filters = make(map[string]Filter)
}

func (s *Storage) SetFilterSymbol(symbol string, filter Filter) {
	s.Locker.Lock()
	defer s.Locker.Unlock()

	s.Filters[symbol] = filter
}

func (s *Storage) SetFilterState(symbol string, state bool) {
	s.Locker.Lock()
	defer s.Locker.Unlock()

	f := s.Filters[symbol]
	f.IsActive = state
	s.Filters[symbol] = f
}

func (s *Storage) SkipFilter(symbol string) bool {
	s.Locker.RLock()
	defer s.Locker.RUnlock()

	var isActiveFilter bool
	for _, symbol := range s.Filters {
		if symbol.IsActive == true {
			isActiveFilter = true
			break
		}
	}

	a := isActiveFilter
	b := s.Filters[symbol].IsActive

	return a && !b
}

func (s *Storage) GetActiveFilterKeys() []string {
	s.Locker.RLock()
	defer s.Locker.RUnlock()

	keys := make([]string, 0)

	for key, filter := range s.Filters {
		if filter.IsActive {
			keys = append(keys, key)
		}
	}

	return keys
}

func (s *Storage) SetRating(key string, rating Rating) {
	s.Locker.Lock()
	defer s.Locker.Unlock()

	s.Symbols[rating.Symbol].Ratings[key] = rating
}

func (s *Storage) SetInfo(key string, info Info) {
	s.Locker.Lock()
	defer s.Locker.Unlock()

	s.Symbols[info.Symbol].Infos[key] = info
}

func (s *Storage) SetSummary(key string, summary Summary) {
	s.Locker.Lock()
	defer s.Locker.Unlock()

	s.Symbols[summary.Symbol].Summaries[key] = summary
}

func (s *Storage) SetHistory(key string, history History) {
	s.Locker.Lock()
	defer s.Locker.Unlock()

	s.Symbols[history.Symbol].Histories[key] = history
}

func (s *Storage) SetMarker(key string, marker Marker) {
	s.Locker.Lock()
	defer s.Locker.Unlock()

	s.Symbols[marker.Symbol].Markers[key] = marker
}

func (s *Storage) SetTrades(key, symbol string, trades []Trade) {
	s.Locker.Lock()
	defer s.Locker.Unlock()

	s.Symbols[symbol].Trades[key] = trades
}

func (s *Storage) GetSync(key string) bool {
	s.Locker.RLock()
	defer s.Locker.RUnlock()

	return s.Sync[key]
}

func (s *Storage) SetSync(key string, state bool) {
	s.Locker.Lock()
	defer s.Locker.Unlock()

	s.Sync[key] = state
}

type RangeFunc func()

func (s *Storage) Range(f RangeFunc) {
	s.Locker.RLock()
	defer s.Locker.RUnlock()

	f()
}

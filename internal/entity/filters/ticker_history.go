package filters

import "time"

type TickerHistoryDiff struct {
	Name     string
	DateFrom time.Time
	DateTo   time.Time
}

type TickerHistoryResult struct {
	Name      string
	PriceFrom string
	PriceTo   string
}

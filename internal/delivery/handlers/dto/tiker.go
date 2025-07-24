package dto

type Ticker struct {
	Name string `json:"ticker"`
}

type TickerParams struct {
	Name     string
	DateFrom string
	DateTo   string
}

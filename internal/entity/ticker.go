package entity

type Ticker struct {
	Name     string
	Price    string
	DateFrom string
	DateTo   string
	Diff     string
}

type TikcerHistory struct {
	Ticker    string
	Price     float64
	Timestamp int64
}

//todo: обудмать как лушче сделать данные

//

package entity

type Ticker struct {
	Ticker string
	Price  float64 // нужна ли вообще ?
}

type TikcerHistory struct {
	Ticker    string
	Price     float64
	Timestamp int64
}

//todo: обудмать как лушче сделать данные

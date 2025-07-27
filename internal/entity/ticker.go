package entity

import "time"

//структура для тикера
type Ticker struct {
	Name  string
	Price string
}

//структура для истории
type TikcerHistory struct {
	Name       string
	Price      string
	Date       time.Time
	DateFrom   string
	DateTo     string
	Difference string
}

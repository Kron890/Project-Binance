package entity

import "time"

// Ticke структура для тикера
type Ticker struct {
	Name  string
	Price string
}

// TikcerHistory структура для истории
type TikcerHistory struct {
	Name       string
	Price      string
	Date       time.Time
	DateFrom   string
	DateTo     string
	Difference string
}

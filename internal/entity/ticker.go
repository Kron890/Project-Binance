package entity

import "time"

type Ticker struct {
	Name     string
	Price    string
	DateFrom string
	DateTo   string
	Diff     string
}

type TikcerHistory struct {
	Name  string
	Price string
	Date  time.Time
}

//todo: обудмать как лушче сделать данные

//

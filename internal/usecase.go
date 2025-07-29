package internal

import "projectBinacne/internal/entity"

type Usecase interface {
	AddTicker(ticker entity.Ticker) error
	FetchTicker(ticker entity.TikcerHistory) (entity.TikcerHistory, error)
	UpdateTickerHistory() error
}

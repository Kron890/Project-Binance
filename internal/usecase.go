package internal

import "projectBinacne/internal/entity"

// Usecase определяет методы для выполнения бизнес-логики
type Usecase interface {
	AddTicker(ticker entity.Ticker) error
	FetchTicker(ticker entity.TikcerHistory) (entity.TikcerHistory, error)
	UpdateTickerHistory() error
	StartProcess()
}

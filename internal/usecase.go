package internal

import "projectBinacne/internal/entity"

type Usecase interface {
	AddTicker(ticker entity.Ticker) error
	FetchTicker(t entity.Ticker) (entity.TikcerHistory, error)
}

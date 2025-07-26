package internal

import "projectBinacne/internal/entity"

type RepoPostgres interface {
	AddTickersList(name string) error
	GetTickersList() ([]entity.Ticker, error)
	FetchTickerHistory(t entity.TikcerHistory) (entity.TikcerHistory, error)
	AddTickersHistory(t entity.TikcerHistory) error
}

type RepoBinance interface {
	GetPrice(ticker string) (string, error)
}

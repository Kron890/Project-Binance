package internal

import (
	"projectBinacne/internal/entity"
	"projectBinacne/internal/entity/filters"
)

type RepoPostgres interface {
	AddTickersList(name string) error
	GetTickersList() ([]entity.Ticker, error)
	FetchTickerHistory(t filters.TickerHistoryDiff) (filters.TickerHistoryResult, error)
	AddTickersHistory(t entity.TikcerHistory) error
}

type RepoBinance interface {
	GetPrice(ticker string) (string, error)
}

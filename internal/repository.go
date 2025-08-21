package internal

import (
	"projectBinacne/internal/entity"
	"projectBinacne/internal/entity/filters"
)

// RepoPostgres определяет методы для взаимодействия с PostgreSQL
type RepoPostgres interface {
	AddTickersList(name string) error
	GetTickersList() ([]entity.Ticker, error)
	FetchTickerHistory(t filters.TickerHistoryDiff) (filters.TickerHistoryResult, error)
	AddTickersHistory(t []entity.TikcerHistory) error
}

// RepoBinance определяет методы для взаимодействия с Binance API
type RepoBinance interface {
	GetPrice(ticker string) (string, error)
	GetPricesList(t []entity.Ticker) ([]entity.TikcerHistory, error)
}

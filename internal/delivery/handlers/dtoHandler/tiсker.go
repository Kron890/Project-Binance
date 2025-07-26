package dto

import (
	"fmt"
	"projectBinacne/internal/entity"
)

type Ticker struct {
	Name string `json:"ticker"`
}

type TickerParams struct {
	Name     string
	DateFrom string
	DateTo   string
}

type TickerResponse struct {
	Name       string `json:"ticker"`
	Price      string `json:"price"`
	Difference string `json:"difference"`
}

func MapTickerToEntity(t Ticker) entity.Ticker {
	return entity.Ticker{Name: t.Name}
}

func MapTickerParamsToEntity(t TickerParams) entity.Ticker {
	return entity.Ticker{
		Name:     t.Name,
		DateFrom: t.DateFrom,
		DateTo:   t.DateTo,
	}
}

func MapEntityToResponce(t entity.Ticker) TickerResponse {
	return TickerResponse{
		Name:       t.Name,
		Price:      t.Price,
		Difference: fmt.Sprintf("%s%%", t.Diff)}
}

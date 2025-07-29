package dto

import (
	"fmt"
	"projectBinacne/internal/entity"
	"strings"
)

type Ticker struct {
	Name string `json:"ticker"`
}

// для запроса
type TickerParams struct {
	Name     string
	DateFrom string
	DateTo   string
}

// для ответа
type TickerResponse struct {
	Name       string `json:"ticker"`
	Price      string `json:"price"`
	Difference string `json:"difference"`
}

func MapTickerToEntity(t Ticker) entity.Ticker {
	return entity.Ticker{Name: t.Name}
}

func MapTickerParamsToHistory(t TickerParams) entity.TikcerHistory {
	return entity.TikcerHistory{
		Name:     t.Name,
		DateFrom: t.DateFrom,
		DateTo:   t.DateTo,
	}
}

func MapEntityToResponce(t entity.TikcerHistory) TickerResponse {
	trimmed := strings.TrimRight(t.Price, "0")
	t.Price = strings.TrimRight(trimmed, ".")

	if t.Difference == "" {
		t.Difference = "0"
	}
	return TickerResponse{
		Name:       t.Name,
		Price:      t.Price,
		Difference: fmt.Sprintf("%s%%", t.Difference)}
}

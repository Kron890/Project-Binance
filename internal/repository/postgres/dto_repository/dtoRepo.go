package dtorepository

import (
	"projectBinacne/internal/entity"
	"time"
)

type TikcerHistory struct {
	Name  string
	Price string
	Date  time.Time
}

func MapEntitesToHistories(tickers []entity.TikcerHistory) []TikcerHistory {
	var result []TikcerHistory
	for _, t := range tickers {
		history := TikcerHistory{
			Name:  t.Name,
			Price: t.Price,
			Date:  t.Date,
		}
		result = append(result, history)
	}
	return result
}

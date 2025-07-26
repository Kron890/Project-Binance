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

func MapEntityToHistory(t entity.TikcerHistory) TikcerHistory {
	return TikcerHistory{
		Name:  t.Name,
		Price: t.Price,
		Date:  t.Date}
}

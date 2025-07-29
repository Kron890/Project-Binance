package dtobinacne

import (
	"projectBinacne/internal/entity"
	"time"

	"github.com/adshao/go-binance/v2"
)

func MapPriceToHistroy(tickers []*binance.SymbolPrice) []entity.TikcerHistory {
	var result []entity.TikcerHistory
	for _, t := range tickers {
		history := entity.TikcerHistory{
			Name:  t.Symbol,
			Price: t.Price,
			Date:  time.Now().Truncate(time.Second),
		}
		result = append(result, history)
	}
	return result
}

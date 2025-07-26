package binanceapi

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
)

type BinanceService struct {
	client           *binance.Client
	listPriceService *binance.ListPricesService
}

func NewBinanceService() *BinanceService {
	client := binance.NewClient("", "")
	service := client.NewListPricesService()

	return &BinanceService{client: client, listPriceService: service}
}

// вытаскивает прайс и отдает в entity.Ticker
func (b *BinanceService) GetPrice(ticker string) (string, error) {
	prices, err := b.listPriceService.Symbol(ticker).Do(context.Background())
	if err != nil {
		return "", err
	}

	if len(prices) == 0 {
		return "", fmt.Errorf("no prices found for symbol: %s", ticker)
	}

	return prices[0].Price, nil

}

// price, err := strconv.ParseFloat(prices[0].Price, 64)
// if err != nil {
// 	return 0, err
// }

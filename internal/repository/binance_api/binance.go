package binanceapi

import (
	"context"
	"fmt"
	"projectBinacne/internal/entity"
	dtobinacne "projectBinacne/internal/repository/binance_api/dtoBinacne"

	"github.com/adshao/go-binance/v2"
)

type BinanceService struct {
	client           *binance.Client
	listPriceService *binance.ListPricesService
}

func NewBinanceService() *BinanceService {
	client := binance.NewClient("", "") //Внешний api binance, здесь не нужен ключ
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

// вытаскиваем несколько тикеров
func (b *BinanceService) GetPricesList(t []entity.Ticker) ([]entity.TikcerHistory, error) {
	tickerList := make([]string, 0, len(t))

	for _, ticker := range t {
		tickerList = append(tickerList, ticker.Name)
	}

	prices, err := b.client.NewListPricesService().Symbols(tickerList).Do(context.Background())
	if err != nil {
		return []entity.TikcerHistory{}, err
	}

	return dtobinacne.MapPriceToHistroy(prices), nil
}

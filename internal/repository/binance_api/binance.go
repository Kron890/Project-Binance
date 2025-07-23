package binanceapi

import "github.com/adshao/go-binance/v2"

type BinanceService struct {
	client           *binance.Client
	listPriceService *binance.ListPricesService
}

func NewBinanceService() *BinanceService {
	client := binance.NewClient("", "")
	service := client.NewListPricesService()

	return &BinanceService{client: client, listPriceService: service}
}

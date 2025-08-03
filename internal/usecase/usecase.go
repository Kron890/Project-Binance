package usecase

import (
	"fmt"
	"log"
	"projectBinacne/internal"
	"projectBinacne/internal/entity"
	"projectBinacne/internal/entity/filters"
	"projectBinacne/pkg"
	"strconv"
	"strings"
	"time"
)

type Ucecase struct {
	Repo           internal.RepoPostgres
	BinanceService internal.RepoBinance
}

func NewUsecase(repo internal.RepoPostgres, binanceService internal.RepoBinance) *Ucecase {
	return &Ucecase{
		Repo:           repo,
		BinanceService: binanceService,
	}
}

func (uc *Ucecase) StartProcess() {
	go uc.startTickerHistoryUpdater()
}

// просто добавляем в бд
func (uc *Ucecase) AddTicker(ticker entity.Ticker) error {

	ticker.Name = strings.ToUpper(ticker.Name)

	_, err := uc.BinanceService.GetPrice(ticker.Name)

	if err != nil {
		log.Print(err)
		return fmt.Errorf("ticker not found")
	}

	err = uc.Repo.AddTickersList(ticker.Name)
	if err != nil {
		return err
	}

	return nil

}

// вытаскиваем данные
func (uc *Ucecase) FetchTicker(ticker entity.TikcerHistory) (entity.TikcerHistory, error) {

	//если нет даты,то вытаскиваем на данный момент
	if ticker.DateFrom == "" || ticker.DateTo == "" {
		price, err := uc.BinanceService.GetPrice(ticker.Name)
		if err != nil {
			return entity.TikcerHistory{}, err
		}

		ticker.Price = price

		return ticker, nil
	}

	dateFrom, dateTo, err := pkg.ParseDate(ticker.DateFrom, ticker.DateTo)
	if err != nil {
		return entity.TikcerHistory{}, err
	}

	history, err := uc.Repo.FetchTickerHistory(filters.TickerHistoryDiff{
		Name:     ticker.Name,
		DateFrom: dateFrom,
		DateTo:   dateTo})
	if err != nil {
		return entity.TikcerHistory{}, err
	}

	ticker.Difference, err = uc.differenceCalculator(history)
	if err != nil {
		return entity.TikcerHistory{}, err
	}

	ticker.Price = history.PriceTo

	return ticker, nil
}

// регулярное обновление данных
func (uc *Ucecase) UpdateTickerHistory() error {
	tickersList, err := uc.Repo.GetTickersList()
	if err != nil {
		return err
	}

	tickersHistory, err := uc.BinanceService.GetPricesList(tickersList)
	if err != nil {
		return err
	}

	return uc.Repo.AddTickersHistory(tickersHistory)

}

func (uc *Ucecase) startTickerHistoryUpdater() {
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		time.Sleep(5 * time.Second)
		for {
			if err := uc.UpdateTickerHistory(); err != nil {
				log.Println("UpdateTickerHistory error:", err)
			}
			<-ticker.C
		}
	}()
}

func (uc *Ucecase) differenceCalculator(result filters.TickerHistoryResult) (string, error) {

	startPrice, err := strconv.ParseFloat(result.PriceFrom, 64)
	if err != nil {
		return "", err
	}
	if startPrice == 0 {
		return "", fmt.Errorf("division by zero")
	}

	endPrice, err := strconv.ParseFloat(result.PriceTo, 64)
	if err != nil {
		return "", err
	}

	diff := ((endPrice - startPrice) / startPrice) * 100
	return fmt.Sprintf("%.2f", diff), nil
}

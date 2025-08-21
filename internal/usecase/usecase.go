package usecase

import (
	"fmt"
	"projectBinacne/internal"
	"projectBinacne/internal/entity"
	"projectBinacne/internal/entity/filters"
	"projectBinacne/pkg"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// Usecase бизнес-логика приложения
type Ucecase struct {
	Repo           internal.RepoPostgres
	BinanceService internal.RepoBinance
	logs           *logrus.Logger
}

// New создаёт слой usecase
func New(repo internal.RepoPostgres, binanceService internal.RepoBinance, logs *logrus.Logger) *Ucecase {
	return &Ucecase{
		Repo:           repo,
		BinanceService: binanceService,
		logs:           logs,
	}
}

// StartProcess запускает фоновые задачи
func (uc *Ucecase) StartProcess() {
	uc.startTickerHistoryUpdater()
}

// AddTicker сохраняет новый тикер в БД
func (uc *Ucecase) AddTicker(ticker entity.Ticker) error {

	ticker.Name = strings.ToUpper(ticker.Name)

	_, err := uc.BinanceService.GetPrice(ticker.Name)

	if err != nil {
		uc.logs.Print(err)
		return fmt.Errorf("ticker not found")
	}

	err = uc.Repo.AddTickersList(ticker.Name)
	if err != nil {
		uc.logs.Error(err)
		return fmt.Errorf("no ticker added")
	}
	return nil

}

// FetchTicker возвращает цену или разницу за период
func (uc *Ucecase) FetchTicker(ticker entity.TikcerHistory) (entity.TikcerHistory, error) {

	//если нет даты,то вытаскиваем данные на данный момент
	if ticker.DateFrom == "" || ticker.DateTo == "" {
		price, err := uc.BinanceService.GetPrice(ticker.Name)
		if err != nil {
			uc.logs.Error(err)
			return entity.TikcerHistory{}, fmt.Errorf("no data received")
		}

		ticker.Price = price

		return ticker, nil
	}

	dateFrom, dateTo, err := pkg.ParseDate(ticker.DateFrom, ticker.DateTo)
	if err != nil {
		uc.logs.Error(err)
		return entity.TikcerHistory{}, fmt.Errorf("no data received")
	}

	history, err := uc.Repo.FetchTickerHistory(filters.TickerHistoryDiff{
		Name:     ticker.Name,
		DateFrom: dateFrom,
		DateTo:   dateTo})
	if err != nil {
		uc.logs.Error(err)
		return entity.TikcerHistory{}, fmt.Errorf("no data received")
	}

	ticker.Difference, err = uc.differenceCalculator(history)
	if err != nil {
		uc.logs.Error(err)
		return entity.TikcerHistory{}, fmt.Errorf("no data received")
	}

	ticker.Price = history.PriceTo

	return ticker, nil
}

// UpdateTickerHistory обновляет исторические цены
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

// startTickerHistoryUpdater запускает фоновый цикл обновления
func (uc *Ucecase) startTickerHistoryUpdater() {
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		time.Sleep(60 * time.Second)
		for {
			if err := uc.UpdateTickerHistory(); err != nil {
				uc.logs.Info("UpdateTickerHistory error:", err)
			}
			<-ticker.C
		}
	}()
}

// differenceCalculator считает процентное изменение цены
func (uc *Ucecase) differenceCalculator(result filters.TickerHistoryResult) (string, error) {

	startPrice, err := strconv.ParseFloat(result.PriceFrom, 64)
	if err != nil {
		uc.logs.Error(err)
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

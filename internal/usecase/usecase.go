package usecase

import (
	"fmt"
	"log"
	"projectBinacne/internal"
	"projectBinacne/internal/entity"
	"projectBinacne/internal/entity/filters"
	"projectBinacne/internal/usecase/helpers"
	"strings"
	"time"
)

type Ucecase struct {
	Repo       internal.RepoPostgres
	BinService internal.RepoBinance
}

func NewUsecase(r internal.RepoPostgres, b internal.RepoBinance) *Ucecase {
	return &Ucecase{
		Repo:       r,
		BinService: b}

}

// просто добавляем в бд
func (uc *Ucecase) AddTicker(ticker entity.Ticker) error {

	ticker.Name = strings.ToUpper(ticker.Name)

	_, err := uc.BinService.GetPrice(ticker.Name)

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
		price, err := uc.BinService.GetPrice(ticker.Name)
		if err != nil {
			return entity.TikcerHistory{}, err
		}

		ticker.Price = price

		return ticker, nil
	}

	dateFrom, dateTo, err := helpers.ParseDate(ticker.DateFrom, ticker.DateTo)
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

	ticker.Difference, err = helpers.DifferenceCalculator(history)
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

	tickersHistory, err := uc.BinService.GetPricesList(tickersList)
	if err != nil {
		return err
	}

	return uc.Repo.AddTickersHistory(tickersHistory)

}

func StartTickerHistoryUpdater(uc *Ucecase) {
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

package usecase

import (
	"fmt"
	"projectBinacne/internal"
	"projectBinacne/internal/entity"
	"time"
)

type Ucecase struct {
	Repo       internal.RepoPostgres
	BinService internal.RepoBinance
}

const format = "02.01.06 15:04:05"

func NewUsecase(r internal.RepoPostgres, b internal.RepoBinance) *Ucecase {
	return &Ucecase{
		Repo:       r,
		BinService: b}

}

// просто добавляем в бд
func (uc *Ucecase) AddTicker(ticker entity.Ticker) error {

	_, err := uc.BinService.GetPrice(ticker.Name)
	if err != nil {
		return err
	}

	err = uc.Repo.AddTickersList(ticker.Name)
	if err != nil {
		return err
	}

	return nil

}

// вытаскиваем данные
func (uc *Ucecase) FetchTicker(t entity.Ticker) (entity.TikcerHistory, error) {

	//если нет даты,то вытаскиваем на данный момент
	if t.DateFrom == "" || t.DateTo == "" {
		price, err := uc.BinService.GetPrice(t.Name)
		if err != nil {
			return entity.TikcerHistory{}, err
		}
		t.Price = price

		return t, nil
	}

	dateFrom, err := time.Parse(t.DateFrom, format)
	if err != nil {
		return entity.TikcerHistory{}, err
	}
	dataFrom := entity.TikcerHistory{Name: t.Name, Date: dateFrom}

	dateTo, err := time.Parse(t.DateTo, format)
	if err != nil {
		return entity.TikcerHistory{}, err
	}

	dataTo := entity.TikcerHistory{Name: t.Name, Date: dateTo}

	dataFrom, err = uc.Repo.FetchTickerHistory(dataFrom)
	if err != nil {
		return entity.TikcerHistory{}, err
	}

	dataTo, err = uc.Repo.FetchTickerHistory(dataTo)
	if err != nil {
		return entity.TikcerHistory{}, err
	}

	fmt.Println(dataFrom, dataTo)

	//считаем разницу и кидаем TikcerHistory

	return entity.TikcerHistory{}, nil
}

//функция которая будет регулярно обновлять данные в бд
// то есть она должна лазить в бд лист
//как-то фиксировать время если в бинансе не приходит со временем
//весь этот лист кидать в бинанс репозиторий, забирать прайс
//конвертировать, записать время и положить в структруру
//отдать в репозиторий и записать в history

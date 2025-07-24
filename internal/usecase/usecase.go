package usecase

import (
	"projectBinacne/internal/entity"
	binanceapi "projectBinacne/internal/repository/binance_api"
	repository "projectBinacne/internal/repository/postgres"
)

type Ucecase struct {
	Repo       *repository.PostgresRepo
	BinService *binanceapi.BinanceService
}

func NewUsecase(r *repository.PostgresRepo, b *binanceapi.BinanceService) *Ucecase {
	return &Ucecase{
		Repo:       r,
		BinService: b}

}

func (uc *Ucecase) AddTicker(ticker entity.Ticker) (entity.Ticker, error) {
	//проверка есть ли такая монета ?

	//кладем в entity

	//отдаем в репозиторий
	return entity.Ticker{}, nil

}

func (uc *Ucecase) FetchTicker(entity.Ticker) (entity.Ticker, error) {
	//берем данные из репо
	//если нет время возвращаем на данный момент
	// и добавляем данные в бд

	//преобразуем время, деньги

	//находим разницу

	//кидаем в структуру и отдаем (все в стринг)

	return entity.Ticker{}, nil
}

//функция которая будет регулярно обновлять данные в бд
// то есть она должна лазить в бд лист
//как-то фиксировать время если в бинансе не приходит со временем
//весь этот лист кидать в бинанс репозиторий, забирать прайс
//конвертировать, записать время и положить в структруру
//отдать в репозиторий и записать в history

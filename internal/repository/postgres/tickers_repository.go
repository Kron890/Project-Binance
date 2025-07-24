package postgres

import (
	"projectBinacne/infrastructure/database"
	"projectBinacne/internal/entity"
)

type PostgresRepo struct {
	DB *database.DataBase
}

func NewRepo(db *database.DataBase) *PostgresRepo {
	return &PostgresRepo{DB: db}
}

//(name string) правильно или лушче в структуру, но там просто один параметр ?

// кладет в бд название тикера(отделаня бд для название тикеров)
func (db *PostgresRepo) AddTickersList(name string) error {
	//смотрим есть ли такой тикер
	//если есть возваращет ответ с ошибкой

	//закидываем тикер
	return nil
}

// вытаскиваем все название
func (db *PostgresRepo) GetTickersList() ([]entity.Ticker, error) {
	return []entity.Ticker{}, nil
}

func (db *PostgresRepo) FetchTickerHistory(t entity.Ticker) (entity.Ticker, error) {
	//map dto
	//вытаскиваем данные из бд
	//отдельно вытаскиваем dateTO dateFROM
	//map Dto
	//отдаем в usecase

	return entity.Ticker{}, nil
}

// кладем данные с историей прайсов
func (db *PostgresRepo) AddTickersHistory(t entity.Ticker) error {
	//mapDto

	//кладем в бд
	return nil
}

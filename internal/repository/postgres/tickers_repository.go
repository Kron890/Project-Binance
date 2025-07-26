package postgres

import (
	"projectBinacne/infrastructure/database"
	"projectBinacne/internal/entity"
	dtorepository "projectBinacne/internal/repository/postgres/dto_repository"
)

type PostgresRepo struct {
	DB *database.DataBase
}

func NewRepo(db *database.DataBase) *PostgresRepo {
	return &PostgresRepo{DB: db}
}

//(name string) правильно или лушче в структуру, но там просто один параметр ?

// кладет в бд название тикера(отделаня бд для название тикеров)
func (r *PostgresRepo) AddTickersList(name string) error {
	//смотрим есть ли такой тикер
	//если есть возваращет ответ с ошибкой

	//закидываем тикер
	return nil
}

// вытаскиваем все название
func (r *PostgresRepo) GetTickersList() ([]entity.Ticker, error) {
	return []entity.Ticker{}, nil
}

// находим данные
func (r *PostgresRepo) FetchTickerHistory(t entity.TikcerHistory) (entity.TikcerHistory, error) {
	//map dto
	// query := "SELECT * FROM ticker_history_list " //запрос

	//парсим в time.time

	// row := r.DB.DB.QueryRow(query, sql.Named("date"))

	//map Dto
	//отдаем в usecase

	return entity.TikcerHistory{}, nil
}

// кладем данные с историей
func (r *PostgresRepo) AddTickersHistory(t entity.TikcerHistory) error {
	ticker := dtorepository.MapEntityToHistory(t)

	query := "INSERT INTO ticker_history_list (ticker, price, date) VALUES ($1, $2, $3)"

	_, err := r.DB.DB.Exec(query, ticker.Name, ticker.Price, ticker.Date)
	if err != nil {
		return err
	}

	return nil
}

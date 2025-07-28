package postgres

import (
	"fmt"
	"projectBinacne/infrastructure/database"
	"projectBinacne/internal/entity"
	"projectBinacne/internal/entity/filters"
	dtorepository "projectBinacne/internal/repository/postgres/dto_repository"
)

type PostgresRepo struct {
	DB *database.DataBase
}

func NewRepo(db *database.DataBase) *PostgresRepo {
	return &PostgresRepo{DB: db}
}

//(name string) правильно или лушче в структуру, но там просто один параметр ?

// кладет в бд название тикера(отделаня бд для название тикеров)(ticker_list)
func (r *PostgresRepo) AddTickersList(ticker string) error {

	query := `INSERT INTO ticker_list (ticker) VALUES ($1) ON CONFLICT DO NOTHING`

	result, err := r.DB.DB.Exec(query, ticker)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("ticker %s already exists", ticker)
	}
	return nil
}

// вытаскиваем все название
func (r *PostgresRepo) GetTickersList() ([]entity.Ticker, error) {
	query := `SELECT ticker FROM ticker_list`

	rows, err := r.DB.DB.Query(query)
	if err != nil {
		return []entity.Ticker{}, err
	}
	defer rows.Close()
	tickerList := make([]entity.Ticker, 0, 10)

	for rows.Next() {
		var t entity.Ticker

		err := rows.Scan(&t.Name)
		if err != nil {
			return []entity.Ticker{}, err
		}

		tickerList = append(tickerList, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tickerList, nil
}

// находим данные
func (r *PostgresRepo) FetchTickerHistory(t filters.TickerHistoryDiff) (filters.TickerHistoryResult, error) {
	result := filters.TickerHistoryResult{Name: t.Name}

	query := `SELECT price FROM ticker_history_list WHERE name = $1 AND date = $2 LIMIT 1`

	// Запрос для даты from
	err := r.DB.DB.QueryRow(query, t.Name, t.DateFrom).Scan(&result.PriceFrom)
	if err != nil {
		return filters.TickerHistoryResult{}, err
	}

	// Запрос для даты to
	err = r.DB.DB.QueryRow(query, t.Name, t.DateTo).Scan(&result.PriceTo)
	if err != nil {
		return filters.TickerHistoryResult{}, err
	}

	return result, nil
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

package postgres

import (
	"fmt"
	"projectBinacne/infrastructure/database"
	"projectBinacne/internal/entity"
	"projectBinacne/internal/entity/filters"
	dtorepository "projectBinacne/internal/repository/postgres/dto_repository"

	"github.com/sirupsen/logrus"
)

// PostgresRepo реализует хранение данных в PostgreSQL
type PostgresRepo struct {
	DB   *database.DataBase
	logs *logrus.Logger
}

// New создаёт новый репозиторий
func New(db *database.DataBase, logs *logrus.Logger) *PostgresRepo {
	return &PostgresRepo{
		DB:   db,
		logs: logs,
	}
}

// AddTickersList сохраняет уникальное название тикера
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
	r.logs.Info("the ticker was successfully added to the database")
	return nil
}

// GetTickersList возвращает все названия тикеров
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

	r.logs.Info("tickers received successfully")
	return tickerList, nil
}

// FetchTickerHistory вытаскивает цену на две даты и считает разницу
func (r *PostgresRepo) FetchTickerHistory(t filters.TickerHistoryDiff) (filters.TickerHistoryResult, error) {
	result := filters.TickerHistoryResult{Name: t.Name}
	query := `SELECT price FROM ticker_history_list WHERE ticker = $1 AND date = $2 LIMIT 1`

	err := r.DB.DB.QueryRow(query, t.Name, t.DateFrom).Scan(&result.PriceFrom)
	if err != nil {
		return filters.TickerHistoryResult{}, err
	}

	err = r.DB.DB.QueryRow(query, t.Name, t.DateTo).Scan(&result.PriceTo)
	if err != nil {
		return filters.TickerHistoryResult{}, err
	}

	r.logs.Info("ticker was found successfully")
	return result, nil
}

// AddTickersHistory пакетно сохраняет исторические цены
func (r *PostgresRepo) AddTickersHistory(t []entity.TikcerHistory) error {
	tickers := dtorepository.MapEntitesToHistories(t)

	query := "INSERT INTO ticker_history_list (ticker, price, date) VALUES ($1, $2, $3)"
	for _, t := range tickers {
		// FIX: Либо батч вставка, либо, на худой конец, транзакция
		_, err := r.DB.DB.Exec(query, t.Name, t.Price, t.Date)
		if err != nil {
			return err
		}

	}
	r.logs.Info("data was successfully added to the database")
	return nil
}

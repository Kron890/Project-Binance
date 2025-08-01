package database

import (
	"database/sql"
	"fmt"
	"projectBinacne/config"

	_ "github.com/lib/pq"
)

type DataBase struct {
	DB *sql.DB
}

func NewConnectDB(c config.Config) (*DataBase, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.User,
		c.Password,
		c.Host,
		c.DBPort,
		c.DB,
	)
	connect, err := sql.Open("postgres", connStr)
	if err != nil {
		return &DataBase{}, err
	}
	err = connect.Ping()
	if err != nil {
		connect.Close()
		return &DataBase{}, err
	}

	tables := `
	CREATE TABLE IF NOT EXISTS ticker_list (
		ticker TEXT PRIMARY KEY
	);
	CREATE TABLE IF NOT EXISTS ticker_history_list (
		ticker    TEXT NOT NULL,
		price     NUMERIC(18,8) NOT NULL,
		date TIMESTAMPTZ NOT NULL 
	);`

	_, err = connect.Exec(tables)
	if err != nil {
		return &DataBase{}, err
	}
	return &DataBase{DB: connect}, nil

}

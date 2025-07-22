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

func NewConnectDB(c config.Config) (DataBase, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DB,
	)
	connect, err := sql.Open("postgres", connStr)
	if err != nil {
		return DataBase{}, err
	}

	return DataBase{DB: connect}, nil

}

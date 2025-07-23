package repository

import "projectBinacne/infrastructure/database"

type PostgresRepo struct {
	DB *database.DataBase
}

func NewRepo(db *database.DataBase) *PostgresRepo {
	return &PostgresRepo{DB: db}
}

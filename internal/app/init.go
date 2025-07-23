package app

import (
	"fmt"
	"projectBinacne/config"
	"projectBinacne/infrastructure/database"
	binanceapi "projectBinacne/internal/repository/binance_api"
	repository "projectBinacne/internal/repository/postgres"
	"projectBinacne/internal/usecase"
)

func Init(srv Server, cfg config.Config) error {

	db, err := database.NewConnectDB(cfg)
	if err != nil {
		return err
	}
	repository := repository.NewRepo(db)

	binanceService := binanceapi.NewBinanceService()

	usecace := usecase.NewUsecase(repository, binanceService)

	fmt.Println(usecace)

	return nil
}

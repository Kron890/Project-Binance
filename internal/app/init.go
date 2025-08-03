package app

import (
	"projectBinacne/config"
	"projectBinacne/infrastructure/database"
	"projectBinacne/internal/delivery/handlers"
	"projectBinacne/internal/delivery/routes"
	binanceapi "projectBinacne/internal/repository/binance_api"
	"projectBinacne/internal/repository/postgres"
	"projectBinacne/internal/usecase"
)

func Init(srv *Server, cfg config.Config) error {

	db, err := database.NewConnectDB(cfg)
	if err != nil {
		return err
	}
	repository := postgres.NewRepo(db)

	binanceService := binanceapi.NewBinanceService()

	uc := usecase.NewUsecase(repository, binanceService)
	uc.StartProcess()

	handler := handlers.NewHandler(uc)

	routes.Init(srv.echo, *handler)

	return nil
}

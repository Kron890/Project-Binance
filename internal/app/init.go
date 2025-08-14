package app

import (
	"projectBinacne/config"
	"projectBinacne/infrastructure/database"
	"projectBinacne/internal/delivery/handlers"
	"projectBinacne/internal/delivery/routes"
	binanceapi "projectBinacne/internal/repository/binance_api"
	"projectBinacne/internal/repository/postgres"
	"projectBinacne/internal/usecase"

	"github.com/sirupsen/logrus"
)

func Init(srv *Server, cfg config.Config, logs *logrus.Logger) error {
	//todo добавить логгирования
	db, err := database.NewConnectDB(cfg)
	if err != nil {
		return err
	}
	repository := postgres.NewRepo(db, logs)

	binanceService := binanceapi.NewBinanceService()

	uc := usecase.NewUsecase(repository, binanceService, logs)
	uc.StartProcess()

	handler := handlers.NewHandler(uc, logs)

	routes.Init(srv.echo, *handler)

	return nil
}

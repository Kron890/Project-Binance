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

// Init собирает все слои приложения и регистрирует маршруты
func Init(srv *Server, cfg config.Config, logs *logrus.Logger) error {

	// подключение к БД
	db, err := database.NewConnectDB(cfg)
	if err != nil {
		return err
	}
	repository := postgres.New(db, logs)

	// внешний клиент к Binance
	binanceService := binanceapi.NewBinanceService()

	// бизнес-логика
	uc := usecase.New(repository, binanceService, logs)
	uc.StartProcess()

	// HTTP-обработчики
	handler := handlers.New(uc, logs)
	routes.Init(srv.echo, *handler)

	return nil
}

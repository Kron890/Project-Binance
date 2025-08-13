package main

import (
	"os"
	"projectBinacne/config"
	"projectBinacne/internal/app"
	"projectBinacne/pkg/logger"
)

func main() {
	logger.Init()

	srv := app.NewServer()

	cfg, err := config.GetConfig()
	if err != nil {
		logger.Log.Error("Error: unable to load configuration:", err)
		os.Exit(1)
	}

	err = app.Init(srv, cfg)
	if err != nil {
		logger.Log.Error("Error: failed to initialize application:", err)
		os.Exit(1)
	}

	err = srv.StartServer(cfg)
	if err != nil {
		logger.Log.Error("Error: server failed to start:", err)
		os.Exit(1)
	}

	logger.Log.Info("Server has successfully shut down")

}

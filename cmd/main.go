package main

import (
	"os"
	"projectBinacne/config"
	"projectBinacne/internal/app"
	"projectBinacne/pkg/logger"
)

func main() {
	logs := logger.Init()

	srv := app.NewServer()

	cfg, err := config.GetConfig()
	if err != nil {
		logs.Error("Error: unable to load configuration:", err)
		os.Exit(1)
	}

	err = app.Init(srv, cfg, logs)
	if err != nil {
		logs.Error("Error: failed to initialize application:", err)
		os.Exit(1)
	}

	err = srv.StartServer(cfg)
	if err != nil {
		logs.Error("Error: server failed to start:", err)
		os.Exit(1)
	}

}

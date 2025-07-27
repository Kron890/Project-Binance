package main

import (
	"fmt"
	"log"
	"os"
	"projectBinacne/config"
	"projectBinacne/internal/app"
)

func main() {
	srv := app.NewServer()

	cfg, err := config.GetConfig()
	if err != nil {
		log.Println("Error: unable to load configuration:", err)
		os.Exit(1)
	}

	err = app.Init(srv, cfg)
	if err != nil {
		log.Println("Error: failed to initialize application:", err)
		os.Exit(1)
	}

	err = srv.StartServer(cfg)
	if err != nil {
		log.Println("Error: server failed to start:", err)
		os.Exit(1)
	}
	fmt.Println("Server started successfully.")

	// TODO: add defer for server shutdown if needed
}

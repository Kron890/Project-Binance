package main

import (
	"fmt"
	"log"
	"os"
	"projectBinacne/config"
	"projectBinacne/internal/app"
)

// TODO:  как ошибки отпаврялть ?

func main() {
	srv := app.NewServer()

	cfg, err := config.GetConfig()
	if err != nil {
		log.Println("failed to load config:", err)
		os.Exit(1)

	}

	err = app.Init(srv, cfg)
	if err != nil {
		log.Println("failed to initialize app:", err)
		os.Exit(1)
	}

	err = srv.StartServer(cfg)
	if err != nil {
		log.Println("failed to start server:", err)
		os.Exit(1)
	}
	fmt.Println("успешно")

	//defer на сервер
}

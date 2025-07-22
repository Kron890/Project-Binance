package app

import (
	"fmt"
	"projectBinacne/config"
	"projectBinacne/infrastructure/database"
)

func Init(srv Server, cfg config.Config) error {

	db, err := database.NewConnectDB(cfg)
	if err != nil {
		return err
	}
	// repository := 

	fmt.Println(db)

	return nil
}

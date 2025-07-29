package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DBPort     string
	User       string
	DB         string
	Password   string
	Host       string
}

func GetConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}
	var c Config
	c.DBPort = os.Getenv("POSTGRES_PORT")
	if c.DBPort == "" {
		return Config{}, errors.New("POSTGRES_PORT not found")
	}
	c.User = os.Getenv("POSTGRES_USER")
	if c.User == "" {
		return Config{}, errors.New("POSTGRES_USER not found")

	}
	c.DB = os.Getenv("POSTGRES_DB")
	if c.DB == "" {
		return Config{}, errors.New("POSTGRES_DB not found")
	}
	c.Password = os.Getenv("POSTGRES_PASSWORD")
	if c.Password == "" {
		return Config{}, errors.New("POSTGRES_PASSWORD not found")
	}
	c.Host = os.Getenv("POSTGRES_HOST")
	if c.Host == "" {
		return Config{}, errors.New("POSTGRES_HOST not found")
	}

	c.ServerPort = os.Getenv("SERVER_PORT")
	if err != nil {
		return Config{}, errors.New("SERVER_PORT not found")
	}
	return c, nil
}

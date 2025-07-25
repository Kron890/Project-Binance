package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	User     string
	DB       string
	Password string
	Host     string
}

//TODO: ПОМЕНЯТЬ ОШИБКИ

func GetConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}
	var c Config
	c.Port = os.Getenv("POSTGRES_PORT")
	if c.Port == "" {
		return Config{}, errors.New("не найден порт")
	}
	c.User = os.Getenv("POSTGRES_USER")
	if c.User == "" {
		return Config{}, errors.New("не найден user")

	}
	c.DB = os.Getenv("POSTGRES_DB")
	if c.DB == "" {
		return Config{}, errors.New("не найден DB")
	}
	c.Password = os.Getenv("POSTGRES_PASSWORD")
	if c.Password == "" {
		return Config{}, errors.New("не найден password")
	}
	c.Host = os.Getenv("POSTGRES_HOST")
	if c.Host == "" {
		return Config{}, errors.New("не найден host")
	}

	return c, nil
}

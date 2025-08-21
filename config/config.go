package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// Config хранит настройки приложения
type Config struct {
	ServerPort string // порт сервера
	DBPort     string // порт БД
	User       string // имя пользователя БД
	DB         string // имя БД
	Password   string // пароль БД
	Host       string // хост БД
}

// GetConfig читает переменные окружения и проверяет их наличие
func GetConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}
	var c Config
	c.DBPort = os.Getenv("POSTGRES_PORT") //TODO: для всех таких проверок лучше писать валидаторы, которые через рефлексию, будут проверять что все четко
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
	if c.ServerPort == "" {
		return Config{}, errors.New("SERVER_PORT not found")
	}
	return c, nil
}

package app

import (
	"fmt"
	"projectBinacne/config"

	"github.com/labstack/echo/v4"
)

// Server оборачивает Echo для запуска HTTP-сервиса
type Server struct {
	echo *echo.Echo
}

// NewServer создаёт новый экземпляр HTTP-сервера
func NewServer() *Server {
	e := echo.New()
	return &Server{echo: e}
}

// StartServer запускает сервер на указанном адресе и порту
func (e *Server) StartServer(c config.Config) error {
	path := fmt.Sprintf("%s:%s", c.Host, c.ServerPort)
	return e.echo.Start(path)
}

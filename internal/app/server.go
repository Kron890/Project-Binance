package app

import (
	"fmt"
	"projectBinacne/config"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	return &Server{echo: e}
}

func (e *Server) StartServer(c config.Config) error {
	path := fmt.Sprintf("%s:%s", c.Host, c.ServerPort)
	return e.echo.Start(path)
}

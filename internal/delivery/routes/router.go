package routes

import (
	"projectBinacne/internal/delivery/handlers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, h handlers.Handler) {
	e.POST("/add_ticker", h.AddTicker)
	e.GET("/fetch/:ticker/:date_from/:date_to", h.FetchTicker)
}

package routes

import (
	"projectBinacne/internal/delivery/handlers"

	"github.com/labstack/echo/v4"
)

// TODO: Добавть delete
// Init регистрирует HTTP-эндпоинты
func Init(e *echo.Echo, h handlers.Handler) {
	e.POST("/add_ticker", h.AddTicker)
	e.GET("/fetch", h.FetchTicker)
}

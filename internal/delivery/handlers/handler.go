package handlers

import (
	"net/http"
	"projectBinacne/internal/delivery/handlers/dto"
	"projectBinacne/internal/usecase"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	uc *usecase.Ucecase
}

func NewHandler(uc *usecase.Ucecase) *Handler {
	return &Handler{uc: uc}
}

//todo сделать функцию на ошибки

func (h Handler) AddTicker(c echo.Context) error {
	var t dto.Ticker

	err := c.Bind(&t)
	if err != nil {
		return err
	}
	if t.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ticker field is required",
		})
	}
	//mapdto
	//usecase
	return nil
}

func (h Handler) FetchTicker(c echo.Context) error {
	var t dto.TickerParams
	t.Name = c.Param("ticker")
	if t.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ticker field is required",
		})
	}

	t.DateFrom = c.Param("date_from")

	t.DateTo = c.Param("date_to")
	//mapdto
	//usecase
	return nil
}

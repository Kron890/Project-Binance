package handlers

import (
	"fmt"
	"net/http"
	"projectBinacne/internal"
	dto "projectBinacne/internal/delivery/handlers/dtoHandler"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	uc internal.Usecase
}

func NewHandler(uc internal.Usecase) *Handler {
	return &Handler{uc: uc}
}

//todo сделать функцию на ошибки

// Кладем данные в бд
func (h Handler) AddTicker(c echo.Context) error {
	var ticker dto.Ticker

	err := c.Bind(&ticker)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if ticker.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ticker field is required",
		})
	}

	err = h.uc.AddTicker(dto.MapTickerToEntity(ticker))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, "ticker added successfully")
}

// поиск прайса и разнцу в цене
func (h Handler) FetchTicker(c echo.Context) error {
	var t dto.TickerParams
	t.Name = c.QueryParam("ticker")
	if t.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ticker field is required",
		})
	}

	t.DateFrom = c.QueryParam("date_from")
	if t.DateFrom == "" {
		fmt.Println(t.DateFrom, "not valid")
		// return c.JSON(http.StatusBadRequest, map[string]string{
		// 	"error": "not valid",
		// })
	}

	t.DateTo = c.QueryParam("date_to")
	if t.DateTo == "" {
		fmt.Println(t.DateTo, "not valid")
		// return c.JSON(http.StatusBadRequest, map[string]string{
		// 	"error": "not valid",
		// })
	}

	ticker, err := h.uc.FetchTicker(dto.MapTickerParamsToHistory(t))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.MapEntityToResponce(ticker))
}

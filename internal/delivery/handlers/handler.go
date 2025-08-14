package handlers

import (
	"net/http"
	"projectBinacne/internal"
	"projectBinacne/internal/delivery/handlers/dto"
	"projectBinacne/pkg/validation"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	uc   internal.Usecase
	logs *logrus.Logger
}

func NewHandler(uc internal.Usecase, logs *logrus.Logger) *Handler {
	return &Handler{
		uc:   uc,
		logs: logs,
	}
}

//TODO: сделать функцию на ошибки

// Кладем данные в бд
func (h Handler) AddTicker(c echo.Context) error {
	var ticker dto.Ticker

	err := c.Bind(&ticker)
	if err != nil {
		h.logs.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "the data cannot be received, check the format of the data being sent",
		})
	}

	err = validation.ValidateStruct(&ticker)
	if err != nil {
		h.logs.Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "the data is not valid",
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
	if !validation.Check(t.Name) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ticker field is required",
		})
	}

	t.DateFrom = c.QueryParam("date_from")

	t.DateTo = c.QueryParam("date_to")

	ticker, err := h.uc.FetchTicker(dto.MapTickerParamsToHistory(t))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.MapEntityToResponce(ticker))
}

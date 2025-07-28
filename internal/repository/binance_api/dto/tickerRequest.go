package dto

import "projectBinacne/internal/entity"

type TickerListRequest struct {
	Name string
}

func MapEntityToRequest(t []entity.Ticker) []TickerListRequest {
	result := make([]TickerListRequest, len(t))
	for i, v := range t {
		result[i] = TickerListRequest{
			Name: v.Name,
		}
	}
	return result
}

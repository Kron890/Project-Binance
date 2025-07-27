package helpers

import (
	"fmt"
	"projectBinacne/internal/entity/filters"
	"strconv"
)

func DifferenceCalculator(result filters.TickerHistoryResult) (string, error) {

	startPrice, err := strconv.ParseFloat(result.PriceFrom, 64)
	if err != nil {
		return "", err
	}
	if startPrice == 0 {
		return "", fmt.Errorf("division by zero")
	}

	endPrice, err := strconv.ParseFloat(result.PriceTo, 64)
	if err != nil {
		return "", err
	}

	diff := ((endPrice - startPrice) / startPrice) * 100
	return fmt.Sprintf("%.2f", diff), nil
}

package helpers

import "time"

const format = "02.01.06 15:04:05"

func ParseDate(dateFromStr string, dateToStr string) (dateFrom time.Time, dateTo time.Time, err error) {
	dateFrom, err = time.Parse(dateFromStr, format)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	dateTo, err = time.Parse(dateToStr, format)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return dateFrom, dateTo, nil
}

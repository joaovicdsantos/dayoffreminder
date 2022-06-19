package utils

import (
	"time"
)

func ValidateDateFormat(date string) (bool, time.Time) {
	parsedDate, err := time.Parse("02-01-2006", date)
	return err == nil, parsedDate
}

func ValidateDateFormatWithFormat(date, format string) (bool, time.Time) {
	parsedDate, err := time.Parse(format, date)
	return err == nil, parsedDate
}

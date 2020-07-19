package utils

import "time"

// EndOfMonth -
func EndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

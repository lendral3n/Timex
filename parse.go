package timex

import (
	"errors"
	"time"
)

const RequestLayout = "2006-01-02"

// ParseRequestDate parses a date string in "YYYY-MM-DD" format to time.Time
func ParseRequestDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, errors.New("empty date string")
	}
	return time.Parse(RequestLayout, dateStr)
}


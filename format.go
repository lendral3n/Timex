package timex

import "time"

const ResponseLayout = "2006-01-02 15:04:05"

// FormatToResponse formats time.Time to string "YYYY-MM-DD HH:mm:ss"
func FormatToResponse(t time.Time) string {
	return t.Format(ResponseLayout)
}


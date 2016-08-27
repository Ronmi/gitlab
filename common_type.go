package gitlab

import "time"

// ParseTime parses time formats in api response
func ParseTime(t string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05-07:00", t)
}

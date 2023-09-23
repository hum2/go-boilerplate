package shared

import "time"

// FormatDateOnly is format time.Time as String.
func FormatDateOnly(t time.Time) string {
	return t.Format(time.DateOnly)
}

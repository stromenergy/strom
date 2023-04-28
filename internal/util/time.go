package util

import "time"

func NewTimeUTC() time.Time {
	return time.Now().UTC()
}

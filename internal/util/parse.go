package util

import (
	"strconv"
	"time"
)

func ParseBool(value string, fallback bool) bool {
	b, err := strconv.ParseBool(value)

	if err != nil {
		b = fallback
	}

	return b
}

func ParseFloat64(str string, fallback float64) float64 {
	f, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return fallback
	}

	return f
}

func ParseInt32(str string, fallback int32) int32 {
	f, err := strconv.ParseInt(str, 10, 32)

	if err != nil {
		return fallback
	}

	return int32(f)
}

func ParseInt64(str string, fallback int64) int64 {
	f, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return fallback
	}

	return f
}

func ParseTime(str string, fallback *time.Time) *time.Time {
	t, err := time.Parse(time.RFC3339, str)

	if err != nil {
		return fallback
	}

	return &t
}
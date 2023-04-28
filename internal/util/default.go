package util

import (
	"database/sql"
	"time"
)

func DefaultBool(i interface{}, fallback bool) bool {
	switch t := i.(type) {
	case sql.NullBool:
		if t.Valid {
			return t.Bool
		}
	case bool:
		return t
	case *bool:
		if t != nil {
			return *t
		}
	}

	return fallback
}

func DefaultFloat(i interface{}, fallback float64) float64 {
	switch t := i.(type) {
	case sql.NullFloat64:
		if t.Valid {
			return t.Float64
		}
	case float64:
		return t
	case *float64:
		if t != nil {
			return *t
		}
	}

	return fallback
}

func DefaultInt(i interface{}, fallback int64) int64 {
	switch t := i.(type) {
	case sql.NullInt64:
		if t.Valid {
			return t.Int64
		}
	case int64:
		return t
	case *int64:
		if t != nil {
			return *t
		}
	}

	return fallback
}

func DefaultString(i interface{}, fallback string) string {
	switch t := i.(type) {
	case sql.NullString:
		if t.Valid {
			return t.String
		}
	case string:
		return t
	case *string:
		if t != nil {
			return *t
		}
	}

	return fallback
}

func DefaultTime(i interface{}, fallback time.Time) time.Time {
	switch t := i.(type) {
	case sql.NullTime:
		if t.Valid {
			return t.Time
		}
	case time.Time:
		return t
	case *time.Time:
		if t != nil {
			return *t
		}
	}

	return fallback
}

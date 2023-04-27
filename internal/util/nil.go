package util

import (
	"database/sql"
	"time"
)

func NilBool(i interface{}) *bool {
	switch t := i.(type) {
	case sql.NullBool:
		if t.Valid {
			return &t.Bool
		}
	case bool:
		return &t
	case *bool:
		return t
	}

	return nil
}

func NilFloat64(i interface{}) *float64 {
	switch t := i.(type) {
	case sql.NullFloat64:
		if t.Valid {
			return &t.Float64
		}
	case float64:
		if t != 0 {
			return &t
		}
	case *float64:
		return t
	}

	return nil
}

func NilInt(i interface{}) *int {
	switch t := i.(type) {
	case sql.NullInt64:
		if t.Valid {
			val := int(t.Int64)
			return &val
		}
	case sql.NullInt32:
		if t.Valid {
			val := int(t.Int32)
			return &val
		}
	case sql.NullInt16:
		if t.Valid {
			val := int(t.Int16)
			return &val
		}
	case int:
		if t != 0 {
			return &t
		}
	case *int:
		return t
	}

	return nil
}

func NilInt16(i interface{}) *int16 {
	switch t := i.(type) {
	case sql.NullInt16:
		if t.Valid {
			return &t.Int16
		}
	case int16:
		if t != 0 {
			return &t
		}
	case *int16:
		return t
	case int:
		if t != 0 {
			val := int16(t)
			return &val
		}
	case *int:
		if t != nil {
			val := int16(*t)
			return &val
		}
	}

	return nil
}

func NilInt32(i interface{}) *int32 {
	switch t := i.(type) {
	case sql.NullInt32:
		if t.Valid {
			return &t.Int32
		}
	case int32:
		if t != 0 {
			return &t
		}
	case *int32:
		return t
	case int:
		if t != 0 {
			val := int32(t)
			return &val
		}
	case *int:
		if t != nil {
			val := int32(*t)
			return &val
		}
	}

	return nil
}

func NilInt64(i interface{}) *int64 {
	switch t := i.(type) {
	case sql.NullInt64:
		if t.Valid {
			return &t.Int64
		}
	case int64:
		if t != 0 {
			return &t
		}
	case *int64:
		return t
	case int:
		if t != 0 {
			val := int64(t)
			return &val
		}
	case *int:
		if t != nil {
			val := int64(*t)
			return &val
		}
	}

	return nil
}

func NilString(i interface{}) *string {
	switch t := i.(type) {
	case sql.NullString:
		if t.Valid {
			return &t.String
		}
	case string:
		if len(t) > 0 {
			return &t
		}
	case *string:
		return t
	}

	return nil
}

func NilTime(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}

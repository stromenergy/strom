package types

import (
	"strings"
	"time"

	"github.com/stromenergy/strom/internal/util"
)

const (
	RFC3339 = "2006-01-02T15:04:05Z"
)

type OcppTime time.Time

func (t OcppTime) MarshalJSON() ([]byte, error) {
	ot := time.Time(t).UTC()
	b := make([]byte, 0, len(RFC3339)+2)
	b = append(b, '"')
	b = ot.AppendFormat(b, RFC3339)
	b = append(b, '"')

	return b, nil
}

func (t OcppTime) String() string {
	return time.Time(t).String()
}

func (t OcppTime) Time() time.Time {
	return (time.Time)(t)
}

func ParseOcppTime(str string, fallback *time.Time) *OcppTime {
	t := util.ParseTime(str, fallback)

	return NilOcppTime(t)
}

func NewOcppTime(t *time.Time) OcppTime {
	ot := NilOcppTime(t)

	if ot == nil {
		return (OcppTime)(util.NewTimeUTC())
	}

	return *ot
}

func NilOcppTime(t *time.Time) *OcppTime {
	if t != nil {
		if t.IsZero() {
			return nil
		}

		ot := (OcppTime)(*t)
		return &ot
	}

	return nil
}

func NilTime(t *OcppTime) *time.Time {
	if t != nil {
		return util.NilTime(t.Time())
	}

	return nil
}

func (t *OcppTime) UnmarshalJSON(data []byte) error {
	dataStr := string(data)

	if dataStr == "null" {
		return nil
	}

	dataStr = strings.Replace(dataStr, `"`, ``, -1)
	parsedTime, err := time.Parse(time.RFC3339, dataStr)

	if err != nil {
		parsedTime, err = time.Parse(RFC3339, dataStr)

		if err != nil {
			return err
		}
	}

	*(*time.Time)(t) = parsedTime
	return nil
}

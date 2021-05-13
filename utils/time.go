package utils

import (
	"database/sql"
	"time"
)

func ToNullableTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{}
	}

	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

func ToRFC3339TimeString(t time.Time) string {
	return t.Format(time.RFC3339)
}

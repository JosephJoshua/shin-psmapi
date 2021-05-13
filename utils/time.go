package utils

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// NullTime is a wrapper around sql.NullTime to allow flattening when json.Marshal is called
type NullTime sql.NullTime

func (t *NullTime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(ToRFC3339TimeString(t.Time))
}

func (t *NullTime) Scan(value interface{}) error {
	v, ok := value.(time.Time)
	if !ok {
		return errors.New("value must be of type time.Time")
	}

	*t = ToNullableTime(v)
	return nil
}

func (t NullTime) Value() (driver.Value, error) {
	return sql.NullTime(t), nil
}

func ToNullableTime(t time.Time) NullTime {
	if t.IsZero() {
		return NullTime{}
	}

	return NullTime{Time: t, Valid: true}
}

func ToRFC3339TimeString(t time.Time) string {
	return t.Format(time.RFC3339)
}

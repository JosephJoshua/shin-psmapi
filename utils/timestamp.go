package utils

import "time"

func ToRFC3339TimeString(t time.Time) string {
	return t.Format(time.RFC3339)
}

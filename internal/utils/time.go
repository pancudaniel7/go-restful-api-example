package utils

import (
	"database/sql"
	"time"
)

// ConvertTimeToNullTime converts a time.Time value to sql.NullTime.
// It sets Valid to true if the time is not the zero value.
func ConvertTimeToNullTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

func ConvertNullTimeToTime(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{}
}

package convert

import (
	"database/sql"
	"time"
)

func SQLStringToPString(value sql.NullString) string {
	if value.Valid {
		return value.String
	}
	return ""
}
func Int64To32(i int64) int32 {
	return int32(i)
}
func SQLTimeToString(t time.Time) string {
	return t.Format(time.DateTime)
}
func SQLNullTimeToString(t sql.NullTime) string {
	if t.Valid {
		return t.Time.Format(time.DateTime)
	}
	return ""
}
func SQLNullInt64ToInt64(i sql.NullInt64) int64 {
	if i.Valid {
		return i.Int64
	}
	return 0
}
func SQLNullFloat64ToFloat64(i sql.NullFloat64) float64 {
	if i.Valid {
		return i.Float64
	}
	return 0
}

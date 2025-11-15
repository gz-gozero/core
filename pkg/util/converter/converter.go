package converter

import (
	"database/sql"
	"time"
)

func ConvertNullString(v sql.NullString) *string {
	if v.Valid {
		return &v.String
	}
	return nil
}

func ConvertNullInt64(v sql.NullInt64) *int64 {
	if v.Valid {
		return &v.Int64
	}
	return nil
}

func ConvertNullFloat64(v sql.NullFloat64) *float64 {
	if v.Valid {
		return &v.Float64
	}
	return nil
}

func ConvertNullTime(v sql.NullTime) *time.Time {
	if v.Valid {
		return &v.Time
	}
	return nil
}

func ConvertStringPtr(v string) *string {
	return &v
}

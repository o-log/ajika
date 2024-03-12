package nullables

import "database/sql"

func NewNullString(value string) sql.NullString {
	return sql.NullString{String: value, Valid: true}
}

func NewNullInt32(value int32) sql.NullInt32 {
	return sql.NullInt32{Int32: value, Valid: true}
}

func NewNullBool(value bool) sql.NullBool {
	return sql.NullBool{Bool: value, Valid: true}
}

func NewNullFloat64(value float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: value, Valid: true}
}

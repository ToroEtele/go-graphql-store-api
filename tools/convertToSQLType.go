package tools

import "database/sql"

func ConvertToNullString(value interface{}) sql.NullString {
	switch v := value.(type) {
	case *string:
		if v == nil {
			return sql.NullString{String: "", Valid: false}
		}
		return sql.NullString{String: *v, Valid: true}
	default:
		return sql.NullString{String: "", Valid: false}
	}
}

func ConvertToNullInt32(value interface{}) sql.NullInt32 {
	switch v := value.(type) {
	case *int:
		if v == nil {
			return sql.NullInt32{Int32: 0, Valid: false}
		}
		return sql.NullInt32{Int32: int32(*v), Valid: true}
	default:
		return sql.NullInt32{Int32: 0, Valid: false}
	}
}

func ConvertToNullFloat64(value interface{}) sql.NullFloat64 {
	switch v := value.(type) {
	case *float64:
		if v == nil {
			return sql.NullFloat64{Float64: 0, Valid: false}
		}
		return sql.NullFloat64{Float64: float64(*v), Valid: true}
	default:
		return sql.NullFloat64{Float64: 0, Valid: false}
	}
}

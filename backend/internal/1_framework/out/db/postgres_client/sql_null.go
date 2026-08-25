package postgres_client

import "database/sql"

func stringFromNullString(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}

	return &value.String
}

func stringToNullString(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  value != "",
	}
}

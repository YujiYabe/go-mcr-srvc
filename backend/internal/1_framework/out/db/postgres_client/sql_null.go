package postgres_client

import "database/sql"

func stringFromNullString(
	value sql.NullString,
) (
	valuePointer *string,
) {
	if !value.Valid {
		return nil
	}

	return &value.String
}

func stringToNullString(
	value string,
) (
	nullString sql.NullString,
) {
	return sql.NullString{
		String: value,
		Valid:  value != "",
	}
}

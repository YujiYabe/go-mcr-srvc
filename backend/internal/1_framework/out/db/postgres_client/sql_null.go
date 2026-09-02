package postgres_client

import "database/sql"

func stringFromNullString(
	value sql.NullString,
) (
	valuePointer *string,
) {
	if !value.Valid {
		valuePointer = nil

		return
	}

	valuePointer = &value.String

	return
}

func stringToNullString(
	value string,
) (
	nullString sql.NullString,
) {
	nullString = sql.NullString{
		String: value,
		Valid:  value != "",
	}

	return
}

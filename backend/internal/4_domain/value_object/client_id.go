package value_object

import (
	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	clientIDLengthMax = 99999999999
	clientIDLengthMin = 0
)

type ClientID struct {
	Content *primitiveObject.PrimitiveString
}

func NewClientID(
	value *string,
) (
	clientID ClientID,
	err error,
) {
	clientID = ClientID{}
	primitiveString := &primitiveObject.PrimitiveString{}

	isNil := primitiveString.CheckNil(value)
	valueString := ""
	if !isNil {
		valueString = *value
	}
	clientID.Content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(clientIDLengthMax),
		primitiveString.WithMinLength(clientIDLengthMin),
	)

	err = clientID.Content.Validation()

	return
}

package value_object

import (
	"backend/internal/4_domain/primitive_object"
)

const (
	clientIDLengthMax = 99999999999
	clientIDLengthMin = 0
)

type ClientID struct {
	Content *primitive_object.PrimitiveString
}

func NewClientID(
	value *string,
) (
	clientID ClientID,
	err error,
) {
	clientID = ClientID{}
	primitiveString := &primitive_object.PrimitiveString{}

	valueString, isNil := primitiveString.CheckNil(value)

	clientID.Content = primitive_object.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(clientIDLengthMax),
		primitiveString.WithMinLength(clientIDLengthMin),
	)

	err = clientID.Content.Validation()

	return
}

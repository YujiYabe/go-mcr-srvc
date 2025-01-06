package value_object

import (
	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	clientIDLengthMax = 99999999999
	clientIDLengthMin = 0
)

type ClientID struct {
	err     error
	content *primitiveObject.PrimitiveString
}

func NewClientID(
	value *string,
) (
	clientID ClientID,
) {
	clientID = ClientID{}
	primitiveString := &primitiveObject.PrimitiveString{}

	isNil := primitiveString.CheckNil(value)
	valueString := ""
	if !isNil {
		valueString = *value
	}
	clientID.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(clientIDLengthMax),
		primitiveString.WithMinLength(clientIDLengthMin),
	)

	if clientID.content.GetError() != nil {
		clientID.SetError(clientID.content.GetError())
	}

	return
}
func (receiver *ClientID) GetError() error {
	return receiver.err
}

func (receiver *ClientID) SetError(
	err error,
) {
	receiver.err = err
}

func (receiver *ClientID) GetValue() string {
	return receiver.content.GetValue()
}

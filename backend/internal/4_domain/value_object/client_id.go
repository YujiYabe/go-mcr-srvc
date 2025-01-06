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
	clientID.SetValue(value)
	return
}

func (receiver *ClientID) SetValue(
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(clientIDLengthMax),
		primitiveString.WithMinLength(clientIDLengthMin),
	)

	if receiver.content.GetError() != nil {
		receiver.SetError(receiver.content.GetError())
	}
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

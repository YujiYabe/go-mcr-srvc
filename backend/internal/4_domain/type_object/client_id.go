package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

var (
	clientIDMaxLength uint = 99
	clientIDMinLength uint
)

type ClientID struct {
	content *primitiveObject.PrimitiveString
}

func NewClientID(
	value *string,
) (
	clientID ClientID,
	err error,
) {
	clientID = ClientID{}
	err = clientID.setValue(value)

	return
}

func (receiver *ClientID) setValue(
	value *string,
) (
	err error,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&clientIDMaxLength),
		primitiveString.WithMinLength(&clientIDMinLength),
	)
	if returnedErr := receiver.content.Validation(); returnedErr != nil {
		err = returnedErr
		return
	}
	err = nil
	return
}

func (receiver ClientID) GetValue() (
	value string,
) {
	value = receiver.content.GetValue()
	return
}

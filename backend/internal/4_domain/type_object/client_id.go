package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

var (
	clientIDMaxLength uint = 99
	clientIDMinLength uint = 0
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
) error {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&clientIDMaxLength),
		primitiveString.WithMinLength(&clientIDMinLength),
	)
	if err := receiver.content.Validation(); err != nil {
		return err
	}
	return nil
}

func (receiver ClientID) GetValue() string {
	return receiver.content.GetValue()
}

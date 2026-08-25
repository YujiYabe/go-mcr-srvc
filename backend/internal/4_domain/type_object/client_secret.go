package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

var (
	clientSecretMaxLength uint = 999
	clientSecretMinLength uint = 0
)

type ClientSecret struct {
	content *primitiveObject.PrimitiveString
}

func NewClientSecret(
	value *string,
) (
	clientSecret ClientSecret,
	err error,
) {
	clientSecret = ClientSecret{}
	err = clientSecret.SetValue(value)

	return
}

func (receiver *ClientSecret) SetValue(
	value *string,
) error {
	primitiveString := &primitiveObject.PrimitiveString{}

	minLength := uint(clientSecretMinLength)
	maxLength := uint(clientSecretMaxLength)

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&maxLength),
		primitiveString.WithMinLength(&minLength),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		return receiver.content.GetError()
	}
	return nil
}

func (receiver *ClientSecret) GetValue() string {
	return receiver.content.GetValue()
}

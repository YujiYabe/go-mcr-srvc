package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

var (
	clientSecretMaxLength uint = 999
	clientSecretMinLength uint
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
	err = clientSecret.setValue(value)

	return
}

func (receiver *ClientSecret) setValue(
	value *string,
) (
	err error,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&clientSecretMaxLength),
		primitiveString.WithMinLength(&clientSecretMinLength),
	)
	if returnedErr := receiver.content.Validation(); returnedErr != nil {
		err = returnedErr

		return
	}
	err = nil

	return
}

func (receiver ClientSecret) GetValue() (
	value string,
) {
	value = receiver.content.GetValue()

	return
}

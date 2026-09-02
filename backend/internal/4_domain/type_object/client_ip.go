package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

const (
	ClientIPHeaderName  primitiveObject.ContextKey = "client-ip"
	ClientIPContextName primitiveObject.ContextKey = "clientIP"
)

var (
	clientIPMaxLength uint = 30
	clientIPMinLength uint
)

type ClientIP struct {
	content *primitiveObject.PrimitiveString
}

func NewClientIP(
	value *string,
) (
	clientIP ClientIP,
	err error,
) {
	clientIP = ClientIP{}
	err = clientIP.setValue(value)

	return
}

func (receiver *ClientIP) setValue(
	value *string,
) (
	err error,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&clientIPMaxLength),
		primitiveString.WithMinLength(&clientIPMinLength),
	)
	if returnedErr := receiver.content.Validation(); returnedErr != nil {
		err = returnedErr

		return
	}
	err = nil

	return
}

func (receiver ClientIP) GetValue() (
	value string,
) {
	value = receiver.content.GetValue()

	return
}

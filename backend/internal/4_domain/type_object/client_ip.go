package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

const (
	ClientIPHeaderName  primitiveObject.ContextKey = "client-ip"
	ClientIPContextName primitiveObject.ContextKey = "clientIP"
)

var (
	clientIPMaxLength uint = 30
	clientIPMinLength uint = 0
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
) error {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&clientIPMaxLength),
		primitiveString.WithMinLength(&clientIPMinLength),
	)
	if err := receiver.content.Validation(); err != nil {
		return err
	}
	return nil
}

func (receiver ClientIP) GetValue() string {
	return receiver.content.GetValue()
}

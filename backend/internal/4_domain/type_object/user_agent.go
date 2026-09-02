package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

const (
	UserAgentHeaderName  primitiveObject.ContextKey = "user-agent"
	UserAgentContextName primitiveObject.ContextKey = "UserAgent"
)

var (
	userAgentMaxLength uint = 50
	userAgentMinLength uint
)

type UserAgent struct {
	content *primitiveObject.PrimitiveString
}

func NewUserAgent(
	value *string,
) (
	userAgent UserAgent,
	err error,
) {
	userAgent = UserAgent{}
	err = userAgent.setValue(value)

	return
}

func (receiver *UserAgent) setValue(
	value *string,
) (
	err error,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&userAgentMaxLength),
		primitiveString.WithMinLength(&userAgentMinLength),
	)
	if returnedErr := receiver.content.Validation(); returnedErr != nil {
		err = returnedErr
		return
	}
	err = nil
	return
}

func (receiver UserAgent) GetValue() (
	value string,
) {
	value = receiver.content.GetValue()
	return
}

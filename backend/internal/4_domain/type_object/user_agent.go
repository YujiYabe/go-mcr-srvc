package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

const (
	UserAgentHeaderName  primitiveObject.ContextKey = "user-agent"
	UserAgentContextName primitiveObject.ContextKey = "UserAgent"
)

var (
	userAgentMaxLength uint = 50
	userAgentMinLength uint = 0
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
	err = userAgent.SetValue(value)

	return
}

func (receiver *UserAgent) SetValue(
	value *string,
) error {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&userAgentMaxLength),
		primitiveString.WithMinLength(&userAgentMinLength),
	)

	if receiver.content.GetError() != nil {
		return receiver.content.GetError()
	}
	return nil
}

func (receiver *UserAgent) GetValue() string {
	return receiver.content.GetValue()
}

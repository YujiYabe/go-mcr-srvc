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
	err = userAgent.setValue(value)

	return
}

func (receiver *UserAgent) setValue(
	value *string,
) error {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&userAgentMaxLength),
		primitiveString.WithMinLength(&userAgentMinLength),
	)
	if err := receiver.content.Validation(); err != nil {
		return err
	}
	return nil
}

func (receiver UserAgent) GetValue() string {
	return receiver.content.GetValue()
}

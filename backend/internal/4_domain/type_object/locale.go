package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

const (
	LocaleHeaderName  primitiveObject.ContextKey = "locale"
	LocaleContextName primitiveObject.ContextKey = "Locale"
)

var (
	localeMaxLength uint = 20
	localeMinLength uint = 0
)

type Locale struct {
	content *primitiveObject.PrimitiveString
}

func NewLocale(
	value *string,
) (
	locale Locale,
	err error,
) {
	locale = Locale{}
	err = locale.SetValue(value)

	return
}

func (receiver *Locale) SetValue(
	value *string,
) error {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&localeMaxLength),
		primitiveString.WithMinLength(&localeMinLength),
	)

	if receiver.content.GetError() != nil {
		return receiver.content.GetError()
	}
	return nil
}

func (receiver *Locale) GetValue() string {
	return receiver.content.GetValue()
}

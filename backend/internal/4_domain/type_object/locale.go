package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

const (
	LocaleHeaderName  primitiveObject.ContextKey = "locale"
	LocaleContextName primitiveObject.ContextKey = "Locale"
)

var (
	localeMaxLength uint = 20
	localeMinLength uint
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
	err = locale.setValue(value)

	return
}

func (receiver *Locale) setValue(
	value *string,
) error {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&localeMaxLength),
		primitiveString.WithMinLength(&localeMinLength),
	)
	if err := receiver.content.Validation(); err != nil {
		return err
	}
	return nil
}

func (receiver Locale) GetValue() string {
	return receiver.content.GetValue()
}

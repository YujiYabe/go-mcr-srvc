package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

const (
	TimeZoneHeaderName  primitiveObject.ContextKey = "time-zone"
	TimeZoneContextName primitiveObject.ContextKey = "TimeZone"
)

var (
	timeZoneMaxLength uint = 9
	timeZoneMinLength uint
)

type TimeZone struct {
	content *primitiveObject.PrimitiveString
}

func NewTimeZone(
	value *string,
) (
	timeZone TimeZone,
	err error,
) {
	timeZone = TimeZone{}
	err = timeZone.setValue(value)

	return
}

func (receiver *TimeZone) setValue(
	value *string,
) (
	err error,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),

		primitiveString.WithMaxLength(&timeZoneMaxLength),
		primitiveString.WithMinLength(&timeZoneMinLength),
	)
	if returnedErr := receiver.content.Validation(); returnedErr != nil {
		err = returnedErr
		return
	}
	err = nil
	return
}

func (receiver TimeZone) GetValue() (
	value string,
) {
	value = receiver.content.GetValue()
	return
}

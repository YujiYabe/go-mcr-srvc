package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

const (
	UserIDHeaderName  primitiveObject.ContextKey = "user-id"
	UserIDContextName primitiveObject.ContextKey = "UserID"
)

var (
	userIDMaxLength uint = 9
	userIDMinLength uint
)

type UserID struct {
	content *primitiveObject.PrimitiveString
}

func NewUserID(
	value *string,
) (
	userID UserID,
	err error,
) {
	userID = UserID{}
	err = userID.setValue(value)

	return
}

func (receiver *UserID) setValue(
	value *string,
) (
	err error,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&userIDMaxLength),
		primitiveString.WithMinLength(&userIDMinLength),
	)
	if returnedErr := receiver.content.Validation(); returnedErr != nil {
		err = returnedErr
		return
	}
	err = nil
	return
}

func (receiver UserID) GetValue() (
	value string,
) {
	value = receiver.content.GetValue()
	return
}

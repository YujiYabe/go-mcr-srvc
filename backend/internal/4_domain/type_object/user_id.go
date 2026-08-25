package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

const (
	UserIDHeaderName  primitiveObject.ContextKey = "user-id"
	UserIDContextName primitiveObject.ContextKey = "UserID"
)

var (
	userIDMaxLength uint = 9
	userIDMinLength uint = 0
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
) error {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&userIDMaxLength),
		primitiveString.WithMinLength(&userIDMinLength),
	)
	if err := receiver.content.Validation(); err != nil {
		return err
	}
	return nil
}

func (receiver UserID) GetValue() string {
	return receiver.content.GetValue()
}

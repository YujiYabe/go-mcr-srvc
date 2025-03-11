package type_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	UserIDHeaderName  primitiveObject.ContextKey = "user-id"
	UserIDContextName primitiveObject.ContextKey = "UserID"
)

var (
	userIDMaxLength uint = 9
	userIDMinLength uint = 0
)

type UserID struct {
	err     error
	content *primitiveObject.PrimitiveString
}

func NewUserID(
	ctx context.Context,
	value *string,
) (
	userID UserID,
) {
	userID = UserID{}
	userID.SetValue(ctx, value)

	return
}

func (receiver *UserID) SetValue(
	ctx context.Context,
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&userIDMaxLength),
		primitiveString.WithMinLength(&userIDMinLength),
	)

	if receiver.content.GetError() != nil {
		receiver.SetError(
			ctx,
			receiver.content.GetError(),
		)
	}
}
func (receiver *UserID) GetError() error {
	return receiver.err
}

func (receiver *UserID) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
}

func (receiver *UserID) GetValue() string {
	return receiver.content.GetValue()
}

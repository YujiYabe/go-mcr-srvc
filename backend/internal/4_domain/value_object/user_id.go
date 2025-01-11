package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
)

const (
	userIDLengthMax = 99999999999
	userIDLengthMin = 0
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
		primitiveString.WithMaxLength(userIDLengthMax),
		primitiveString.WithMinLength(userIDLengthMin),
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
	pkg.Logging(ctx, receiver.GetError())
}

func (receiver *UserID) GetValue() string {
	return receiver.content.GetValue()
}

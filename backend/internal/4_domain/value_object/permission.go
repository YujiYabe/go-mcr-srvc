package value_object

import (
	"context"
	"fmt"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
)

const (
	permissionLengthMax = 50
	permissionLengthMin = 1
)

var permissionCheckSpell = []string{}

type Permission struct {
	err     error
	content *primitiveObject.PrimitiveString
}

func NewPermission(
	ctx context.Context,
	value *string,
) (
	permission Permission,
) {
	permission = Permission{}
	permission.SetValue(ctx, value)

	return
}

func (receiver *Permission) SetValue(
	ctx context.Context,
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(permissionLengthMax),
		primitiveString.WithMinLength(permissionLengthMin),
		primitiveString.WithCheckSpell(permissionCheckSpell),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(ctx, receiver.content.GetError())
		return
	}

}

func (receiver *Permission) GetValue() string {
	return receiver.content.GetValue()
}

func (receiver *Permission) GetError() error {
	return receiver.err
}

func (receiver *Permission) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.GetError())
}

func (receiver *Permission) SetErrorString(
	ctx context.Context,
	errString string,
) {
	receiver.SetError(
		ctx,
		fmt.Errorf(
			"error: %s",
			errString,
		),
	)
}

func (receiver *Permission) GetIsNil() bool {
	return receiver.content.GetIsNil()
}

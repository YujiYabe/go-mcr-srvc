package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
)

const (
	PermissionMetaName    primitiveObject.ContextKey = "permission"
	PermissionContextName primitiveObject.ContextKey = "Permission"
)

const (
	permissionLengthMax = 99999999999
	permissionLengthMin = 0
)

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
	// 値の格納前にバリデーション。
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(permissionLengthMax),
		primitiveString.WithMinLength(permissionLengthMin),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(
			ctx,
			receiver.content.GetError(),
		)
	}
}

func (receiver *Permission) GetError() error {
	return receiver.err
}

func (receiver *Permission) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.err)
}

func (receiver *Permission) GetValue() string {
	return receiver.content.GetValue()
}

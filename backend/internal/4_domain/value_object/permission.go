package value_object

import (
	primitiveObject "backend/internal/4_domain/primitive_object"
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
	value *string,
) (
	permission Permission,
) {
	permission = Permission{}
	permission.SetValue(value)
	return
}

func (receiver *Permission) SetValue(
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
		receiver.SetError(receiver.content.GetError())
	}
}

func (receiver *Permission) GetError() error {
	return receiver.err
}

func (receiver *Permission) SetError(
	err error,
) {
	receiver.err = err
}

func (receiver *Permission) GetValue() string {
	return receiver.content.GetValue()
}

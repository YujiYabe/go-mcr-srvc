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
	primitiveString := &primitiveObject.PrimitiveString{}

	isNil := primitiveString.CheckNil(value)
	valueString := ""
	if !isNil {
		valueString = *value
	}
	permission.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(permissionLengthMax),
		primitiveString.WithMinLength(permissionLengthMin),
	)

	permission.content.Validation()
	if permission.content.GetError() != nil {
		permission.SetError(permission.content.GetError())
	}

	return
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

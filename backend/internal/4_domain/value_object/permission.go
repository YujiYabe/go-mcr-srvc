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
	Err     error
	Content *primitiveObject.PrimitiveString
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
	permission.Content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(permissionLengthMax),
		primitiveString.WithMinLength(permissionLengthMin),
	)

	permission.Content.Validation()
	if permission.Content.GetError() != nil {
		permission.SetError(permission.Content.GetError())
	}

	return
}
func (receiver *Permission) GetError() error {
	return receiver.Err
}

func (receiver *Permission) SetError(
	err error,
) {
	receiver.Err = err
}

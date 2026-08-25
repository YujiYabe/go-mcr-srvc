package type_object

import (
	"fmt"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

var (
	permissionMaxLength uint = 50
	permissionMinLength uint = 1
)

var permissionCheckSpell = []string{}

type Permission struct {
	content *primitiveObject.PrimitiveString
}

func NewPermission(
	value *string,
) (
	permission Permission,
	err error,
) {
	permission = Permission{}
	err = permission.SetValue(value)

	return
}

func (receiver *Permission) SetValue(
	value *string,
) error {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&permissionMaxLength),
		primitiveString.WithMinLength(&permissionMinLength),
		primitiveString.WithCheckSpell(permissionCheckSpell),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		return receiver.content.GetError()
	}
	return nil
}

func (receiver *Permission) GetValue() string {
	return receiver.content.GetValue()
}

func (receiver *Permission) ErrorString(
	errString string,
) error {
	return fmt.Errorf("error: %s", errString)
}

func (receiver *Permission) GetIsNil() bool {
	return receiver.content.GetIsNil()
}

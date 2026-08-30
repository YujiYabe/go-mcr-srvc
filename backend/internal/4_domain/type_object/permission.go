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
	err = permission.setValue(value)

	return
}

func (receiver *Permission) setValue(
	value *string,
) (
	err error,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&permissionMaxLength),
		primitiveString.WithMinLength(&permissionMinLength),
		primitiveString.WithCheckSpell(permissionCheckSpell),
	)
	if err := receiver.content.Validation(); err != nil {
		return err
	}
	return nil
}

func (receiver Permission) GetValue() (
	value string,
) {
	return receiver.content.GetValue()
}

func (receiver *Permission) ErrorString(
	errString string,
) (
	err error,
) {
	return fmt.Errorf("error: %s", errString)
}

func (receiver Permission) GetIsNil() (
	ok bool,
) {
	return receiver.content.GetIsNil()
}

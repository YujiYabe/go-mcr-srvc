package type_object

import (
	"fmt"
	"slices"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

var (
// permissionListMaxLength = 50
// permissionListMinLength = 1
)

const (
	PermissionListHeaderName  primitiveObject.ContextKey = "permissions"
	PermissionListContextName primitiveObject.ContextKey = "permissionList"

	PermissionUserRead  = "user:read"
	PermissionUserWrite = "user:write"
)

type PermissionList struct {
	content []Permission
}

func NewPermissionList(
	valueList []string,
) (
	permissionList PermissionList,
	err error,
) {
	permissionList = PermissionList{}
	err = permissionList.setValue(valueList)

	return
}

func (receiver *PermissionList) setValue(
	valueList []string,
) (
	err error,
) {

	for _, value := range valueList {
		permission, err := NewPermission(&value)

		if err != nil {
			return err
		}
		receiver.content = append(
			receiver.content,
			permission,
		)
	}
	return nil
}

func (receiver *PermissionList) ErrorString(
	errString string,
) (
	err error,
) {
	return fmt.Errorf("error: %s", errString)
}

func (receiver PermissionList) GetSliceValue() (
	sliceValue []string,
) {
	sliceValue = []string{}

	for _, permission := range receiver.content {
		sliceValue = append(
			sliceValue,
			permission.GetValue(),
		)
	}

	return sliceValue
}

func (receiver PermissionList) IsEmpty() (
	isEmpty bool,
) {
	return len(receiver.content) == 0
}

func (receiver PermissionList) Count() (
	value int,
) {
	return len(receiver.content)
}

func (receiver PermissionList) Has(
	permission Permission,
) (
	has bool,
) {
	return slices.Contains(receiver.GetSliceValue(), permission.GetValue())
}

func (receiver PermissionList) HasValue(
	value string,
) (
	hasValue bool,
) {
	permission, err := NewPermission(&value)
	if err != nil {
		return false
	}

	return receiver.Has(permission)
}

func (receiver PermissionList) CanReadUser() (
	canReadUser bool,
) {
	return receiver.HasValue(PermissionUserRead)
}

func (receiver PermissionList) CanWriteUser() (
	canWriteUser bool,
) {
	return receiver.HasValue(PermissionUserWrite)
}

func (receiver PermissionList) EnsureHas(
	value string,
) (
	err error,
) {
	if receiver.HasValue(value) {
		return nil
	}

	return receiver.ErrorString(fmt.Sprintf("permission is required: %s", value))
}

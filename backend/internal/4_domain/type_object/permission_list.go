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
		permission, returnedErr := NewPermission(&value)

		if returnedErr != nil {
			err = returnedErr
			return
		}
		receiver.content = append(
			receiver.content,
			permission,
		)
	}
	err = nil
	return
}

func (receiver *PermissionList) ErrorString(
	errString string,
) (
	err error,
) {
	err = fmt.Errorf("error: %s", errString)
	return
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

	return
}

func (receiver PermissionList) IsEmpty() (
	isEmpty bool,
) {
	isEmpty = len(receiver.content) == 0
	return
}

func (receiver PermissionList) Count() (
	value int,
) {
	value = len(receiver.content)
	return
}

func (receiver PermissionList) Has(
	permission Permission,
) (
	has bool,
) {
	has = slices.Contains(receiver.GetSliceValue(), permission.GetValue())
	return
}

func (receiver PermissionList) HasValue(
	value string,
) (
	hasValue bool,
) {
	permission, err := NewPermission(&value)
	if err != nil {
		hasValue = false
		return
	}

	hasValue = receiver.Has(permission)
	return
}

func (receiver PermissionList) CanReadUser() (
	canReadUser bool,
) {
	canReadUser = receiver.HasValue(PermissionUserRead)
	return
}

func (receiver PermissionList) CanWriteUser() (
	canWriteUser bool,
) {
	canWriteUser = receiver.HasValue(PermissionUserWrite)
	return
}

func (receiver PermissionList) EnsureHas(
	value string,
) (
	err error,
) {
	if receiver.HasValue(value) {
		err = nil
		return
	}

	err = receiver.ErrorString(fmt.Sprintf("permission is required: %s", value))
	return
}

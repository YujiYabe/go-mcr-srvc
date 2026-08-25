package type_object

import (
	"fmt"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

var (
// permissionListMaxLength = 50
// permissionListMinLength = 1
)

const (
	PermissionListHeaderName  primitiveObject.ContextKey = "permissions"
	PermissionListContextName primitiveObject.ContextKey = "permissionList"
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
	err = permissionList.SetValue(valueList)

	return
}

func (receiver *PermissionList) SetValue(
	valueList []string,
) error {

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
) error {
	return fmt.Errorf("error: %s", errString)
}

func (receiver *PermissionList) GetSliceValue() (
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

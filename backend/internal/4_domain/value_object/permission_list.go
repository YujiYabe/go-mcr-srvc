package value_object

import (
	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	PermissionListMetaName    primitiveObject.ContextKey = "permission-list"
	PermissionListContextName primitiveObject.ContextKey = "PermissionList"
)

const (
// permissionListLengthMax = 99999999999
// permissionListLengthMin = 0
)

type PermissionList struct {
	Err     error
	Content []Permission
}

func NewPermissionList(
	valueList []*string,
) (
	permissionList PermissionList,
) {

	permissionList = PermissionList{
		Content: []Permission{},
	}

	for _, value := range valueList {
		permission := NewPermission(value)
		if permission.Err != nil {
			permissionList.SetError(permission.Err)
			return
		}

		permissionList.Content = append(permissionList.Content, permission)
	}

	return
}

func (receiver *PermissionList) GetError() error {
	return receiver.Err
}

func (receiver *PermissionList) SetError(
	err error,
) {
	if receiver.Err == nil {
		receiver.Err = err
	}
}

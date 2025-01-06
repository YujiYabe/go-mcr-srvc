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
	err     error
	content []Permission
}

func NewPermissionList(
	valueList []*string,
) (
	permissionList PermissionList,
) {

	permissionList = PermissionList{
		content: []Permission{},
	}

	for _, value := range valueList {
		permission := NewPermission(value)
		if permission.GetError() != nil {
			permissionList.SetError(permission.GetError())
			return
		}

		permissionList.content = append(
			permissionList.content,
			permission,
		)
	}

	return
}

func (receiver *PermissionList) GetError() error {
	return receiver.err
}

func (receiver *PermissionList) SetError(
	err error,
) {
	if receiver.err == nil {
		receiver.err = err
	}
}

package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
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
	ctx context.Context,
	valueList []*string,
) (
	permissionList PermissionList,
) {

	permissionList = PermissionList{
		content: []Permission{},
	}

	for _, value := range valueList {
		permission := NewPermission(ctx, value)
		if permission.GetError() != nil {
			permissionList.SetError(
				ctx,
				permission.GetError(),
			)
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
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.err)
}

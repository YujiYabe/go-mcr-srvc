package value_object

import (
	"context"
	"fmt"

	"backend/pkg"
)

const (
// permissionListLengthMax = 50
// permissionListLengthMin = 1
)

type PermissionList struct {
	err     error
	content []Permission
}

func NewPermissionList(
	ctx context.Context,
	valueList []string,
) (
	permissionList PermissionList,
) {
	permissionList = PermissionList{}
	permissionList.SetValue(ctx, valueList)

	return
}

func (receiver *PermissionList) SetValue(
	ctx context.Context,
	valueList []string,
) {

	for _, value := range valueList {
		permission := NewPermission(ctx, &value)

		if permission.GetError() != nil {
			receiver.SetError(ctx, permission.GetError())
			break
		}
		receiver.content = append(
			receiver.content,
			permission,
		)
	}
}

// func (receiver *PermissionList) GetValue() string {
// 	return receiver.content.GetValue()
// }

func (receiver *PermissionList) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.GetError())
}

func (receiver *PermissionList) GetError() error {
	return receiver.err
}

func (receiver *PermissionList) SetErrorString(
	ctx context.Context,
	errString string,
) {
	receiver.SetError(
		ctx,
		fmt.Errorf(
			"error: %s",
			errString,
		),
	)
}

// func (receiver *PermissionList) GetIsNil() bool {
// 	return receiver.content.GetIsNil()
// }

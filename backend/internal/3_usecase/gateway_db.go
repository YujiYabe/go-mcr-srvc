package usecase

import (
	"context"
	"fmt"

	groupObject "backend/internal/4_domain/group_object"
)

type (

	// ToGatewayDB ...
	ToGatewayDB interface {
		RunInTransaction(
			ctx context.Context,
			fn func(context.Context) error,
		) error

		GetUserList(
			ctx context.Context,
		) (
			userList groupObject.UserList,
			err error,
		)

		GetUserListByCondition(
			ctx context.Context,
			reqUser groupObject.User,
		) (
			resUserList groupObject.UserList,
			err error,
		)

		UpdateUser(
			ctx context.Context,
			newUser groupObject.User,
		) error

		UpdateUserEmployment(
			ctx context.Context,
			userEmployment groupObject.UserEmployment,
		) error
	}
)

func (receiver *useCase) GetUserList(
	ctx context.Context,
) (
	userList groupObject.UserList,
	err error,
) {
	if err = ensureContextReady(ctx, "GetUserList"); err != nil {
		return
	}
	userList, err = receiver.ToGatewayDB.GetUserList(ctx)
	if err != nil {
		err = fmt.Errorf("GetUserList: %w", err)
	}
	return
}

func (receiver *useCase) GetUserListByCondition(
	ctx context.Context,
	reqUser groupObject.User,
) (
	resUserList groupObject.UserList,
	err error,
) {
	if err = ensureContextReady(ctx, "GetUserListByCondition"); err != nil {
		return
	}
	if !reqUser.CanBeUsedAsSearchCondition() {
		err = fmt.Errorf("GetUserListByCondition: user search condition is required")
		return
	}

	resUserList, err = receiver.ToGatewayDB.GetUserListByCondition(
		ctx,
		reqUser,
	)
	if err != nil {
		err = fmt.Errorf("GetUserListByCondition: %w", err)
	}
	return
}

func (receiver *useCase) UpdateUser(
	ctx context.Context,
	newUser groupObject.User,
) error {
	if err := ensureContextReady(ctx, "UpdateUser"); err != nil {
		return err
	}
	if err := newUser.EnsureReadyToUpdate(); err != nil {
		return fmt.Errorf("UpdateUser: %w", err)
	}

	if err := receiver.ToGatewayDB.RunInTransaction(
		ctx,
		func(txCtx context.Context) error {
			return receiver.ToGatewayDB.UpdateUser(txCtx, newUser)
		},
	); err != nil {
		return fmt.Errorf("UpdateUser: %w", err)
	}

	return nil
}

func (receiver *useCase) UpdateUserProfileWithPrimaryEmployment(
	ctx context.Context,
	newUser groupObject.User,
	userEmployment groupObject.UserEmployment,
) error {
	if err := ensureContextReady(ctx, "UpdateUserProfileWithPrimaryEmployment"); err != nil {
		return err
	}
	if err := newUser.EnsureReadyToUpdate(); err != nil {
		return fmt.Errorf("UpdateUserProfileWithPrimaryEmployment: %w", err)
	}
	if err := receiver.ToDomain.EnsurePrimaryEmploymentAssignable(newUser, userEmployment); err != nil {
		return fmt.Errorf("UpdateUserProfileWithPrimaryEmployment: %w", err)
	}

	if err := receiver.ToGatewayDB.RunInTransaction(
		ctx,
		func(txCtx context.Context) error {
			if err := receiver.ToGatewayDB.UpdateUser(txCtx, newUser); err != nil {
				return err
			}
			return receiver.ToGatewayDB.UpdateUserEmployment(txCtx, userEmployment)
		},
	); err != nil {
		return fmt.Errorf("UpdateUserProfileWithPrimaryEmployment: %w", err)
	}

	return nil
}

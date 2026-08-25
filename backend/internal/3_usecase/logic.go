package usecase

import (
	"context"
	"fmt"

	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
)

// Start ...
func (receiver *useCase) Start() {
}

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

func (receiver *useCase) FetchAccessToken(
	ctx context.Context,
	credential groupObject.Credential,
) (
	accessToken typeObject.AccessToken,
	err error,
) {
	if err = ensureContextReady(ctx, "FetchAccessToken"); err != nil {
		return
	}
	if err = credential.EnsureReadyToAuthenticate(); err != nil {
		err = fmt.Errorf("FetchAccessToken: %w", err)
		return
	}
	accessToken, err = receiver.ToGatewayExternal.FetchAccessToken(
		ctx,
		credential,
	)
	if err != nil {
		err = fmt.Errorf("FetchAccessToken: %w", err)
	}
	return
}

func (receiver *useCase) ViaGRPC(
	ctx context.Context,
	reqUser groupObject.User,
) (
	resUserList groupObject.UserList,
	err error,
) {
	if err = ensureContextReady(ctx, "ViaGRPC"); err != nil {
		return
	}
	if !reqUser.CanBeUsedAsSearchCondition() {
		err = fmt.Errorf("ViaGRPC: user search condition is required")
		return
	}
	resUserList, err = receiver.ToGatewayExternal.ViaGRPC(
		ctx,
		reqUser,
	)
	if err != nil {
		err = fmt.Errorf("ViaGRPC: %w", err)
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

func (receiver *useCase) UpdateUserWithEmployment(
	ctx context.Context,
	newUser groupObject.User,
	userEmployment groupObject.UserEmployment,
) error {
	if err := ensureContextReady(ctx, "UpdateUserWithEmployment"); err != nil {
		return err
	}
	if err := newUser.EnsureReadyToUpdate(); err != nil {
		return fmt.Errorf("UpdateUserWithEmployment: %w", err)
	}
	if err := userEmployment.EnsureBelongsTo(newUser); err != nil {
		return fmt.Errorf("UpdateUserWithEmployment: %w", err)
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
		return fmt.Errorf("UpdateUserWithEmployment: %w", err)
	}

	return nil
}

func (receiver *useCase) PublishTestTopic(
	ctx context.Context,
) error {
	if err := ensureContextReady(ctx, "PublishTestTopic"); err != nil {
		return err
	}
	if err := receiver.ToGatewayExternal.PublishTestTopic(ctx); err != nil {
		return fmt.Errorf("PublishTestTopic: %w", err)
	}
	return nil
}

func ensureContextReady(
	ctx context.Context,
	usecaseName string,
) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("%s: context is not ready: %w", usecaseName, err)
	}

	return nil
}

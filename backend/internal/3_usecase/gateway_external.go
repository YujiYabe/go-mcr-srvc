package usecase

import (
	"context"
	"fmt"

	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
)

type (

	// ToGatewayExternal ...
	ToGatewayExternal interface {
		FetchAccessToken(
			ctx context.Context,
			credential groupObject.Credential,
		) (
			accessToken typeObject.AccessToken,
			err error,
		)

		ViaGRPC(
			ctx context.Context,
			reqUser groupObject.User,
		) (
			resUserList groupObject.UserList,
			err error,
		)

		PublishTestTopic(
			ctx context.Context,
		) error
	}
)

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

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

func (receiver *useCase) GetPersonList(
	ctx context.Context,
) (
	personList groupObject.PersonList,
	err error,
) {
	if err = ensureContextReady(ctx, "GetPersonList"); err != nil {
		return
	}
	personList, err = receiver.ToGatewayDB.GetPersonList(ctx)
	if err != nil {
		err = fmt.Errorf("GetPersonList: %w", err)
	}
	return
}

func (receiver *useCase) GetPersonListByCondition(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
	err error,
) {
	if err = ensureContextReady(ctx, "GetPersonListByCondition"); err != nil {
		return
	}

	resPersonList, err = receiver.ToGatewayDB.GetPersonListByCondition(
		ctx,
		reqPerson,
	)
	if err != nil {
		err = fmt.Errorf("GetPersonListByCondition: %w", err)
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
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
	err error,
) {
	if err = ensureContextReady(ctx, "ViaGRPC"); err != nil {
		return
	}
	resPersonList, err = receiver.ToGatewayExternal.ViaGRPC(
		ctx,
		reqPerson,
	)
	if err != nil {
		err = fmt.Errorf("ViaGRPC: %w", err)
	}
	return
}

func (receiver *useCase) PublishTestTopic(
	ctx context.Context,
) {
	if err := ensureContextReady(ctx, "PublishTestTopic"); err != nil {
		return
	}
	receiver.ToGatewayExternal.PublishTestTopic(ctx)
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

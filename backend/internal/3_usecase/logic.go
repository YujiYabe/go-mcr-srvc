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
	if !reqPerson.CanBeUsedAsSearchCondition() {
		err = fmt.Errorf("GetPersonListByCondition: person search condition is required")
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
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
	err error,
) {
	if err = ensureContextReady(ctx, "ViaGRPC"); err != nil {
		return
	}
	if !reqPerson.CanBeUsedAsSearchCondition() {
		err = fmt.Errorf("ViaGRPC: person search condition is required")
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

func (receiver *useCase) UpdatePerson(
	ctx context.Context,
	newPerson groupObject.Person,
) error {
	if err := ensureContextReady(ctx, "UpdatePerson"); err != nil {
		return err
	}
	if err := newPerson.EnsureReadyToUpdate(); err != nil {
		return fmt.Errorf("UpdatePerson: %w", err)
	}

	if err := receiver.ToGatewayDB.RunInTransaction(
		ctx,
		func(txCtx context.Context) error {
			return receiver.ToGatewayDB.UpdatePerson(txCtx, newPerson)
		},
	); err != nil {
		return fmt.Errorf("UpdatePerson: %w", err)
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

package db_gateway

import (
	"context"
	"fmt"

	groupObject "backend/internal/4_domain/group_object"
	"backend/internal/logger"
)

// GetPersonList ...
func (receiver *GatewayDB) GetPersonList(
	ctx context.Context,

) (
	personList groupObject.PersonList,
) {
	return receiver.ToPostgres.GetPersonList(ctx)
}

// GetPersonListByCondition ...
func (receiver *GatewayDB) GetPersonListByCondition(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
) {
	resPersonList = receiver.ToPostgres.GetPersonListByCondition(
		ctx,
		reqPerson,
	)

	return
}

// UpdateProduct ...
func (receiver *GatewayDB) UpdateProduct(
	ctx context.Context,
	newProduct groupObject.Product,
) (
	err error,
) {
	isSuccess := true
	tx := receiver.ToPostgres.BeginTx(ctx)
	defer func() {
		if ctx.Err() != nil {
			endTxErr := receiver.ToPostgres.EndTx(ctx, tx, false)
			err = fmt.Errorf("context canceled (%w); rollback result: %v", ctx.Err(), endTxErr)
			return
		}
		endTxErr := receiver.ToPostgres.EndTx(ctx, tx, isSuccess)
		if endTxErr != nil && err == nil {
			err = endTxErr
		}

	}()

	var product domain.Product

	if isSuccess {
		product, err = receiver.ToPostgres.GetProduct(
			tx,
			newProduct.JANCode,
		)
		if err != nil {
			isSuccess = false
			logger.Logging(ctx, err)
		}
	}

	if isSuccess {
		err = receiver.ToPostgres.UpdateProduct(
			tx,
			product,
			newProduct,
		)
		if err != nil {
			isSuccess = false
			logger.Logging(ctx, err)
		}
	}

	return err
}

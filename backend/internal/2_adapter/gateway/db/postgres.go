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
	return receiver.ToPostgres.GetPersonList(
		ctx,
		receiver.ToPostgres.WithOutTx(ctx),
	)
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
		receiver.ToPostgres.WithOutTx(ctx),
		reqPerson,
	)

	return
}

// UpdatePerson ...
func (receiver *GatewayDB) UpdatePerson(
	ctx context.Context,
	newPerson groupObject.Person,
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

	var person groupObject.Person

	if isSuccess {
		person = receiver.ToPostgres.GetPerson(
			ctx,
			tx,
			newPerson.ID,
		)
		if person.GetError() != nil {
			isSuccess = false
			logger.Logging(ctx, err)
		}
	}

	// if isSuccess {
	// 	err = receiver.ToPostgres.UpdatePerson(
	// 		tx,
	// 		person,
	// 		newPerson,
	// 	)
	// 	if err != nil {
	// 		isSuccess = false
	// 		logger.Logging(ctx, err)
	// 	}
	// }

	return err
}

package db_gateway

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
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

// // UpdateProduct ...
// func (receiver *GatewayDB) UpdateProduct(
// 	ctx context.Context,
// 	newProduct groupObject.Product,
// ) (
// 	err error,
// ) {
// 	isSuccess := true
// 	tx := receiver.ToPostgres.BeginTx()
// 	defer func() {
// 		receiver.ToPostgres.EndTx(tx, isSuccess)
// 	}()

// 	var product domain.Product

// 	if isSuccess {
// 		var err error
// 		product, err = receiver.ToPostgres.GetProduct(
// 			tx,
// 			newProduct.JANCode,
// 		)
// 		if err != nil {
// 			isSuccess = false
// 			logger.Logging(ctx, err)
// 		}
// 	}

// 	if isSuccess {
// 		err = receiver.ToPostgres.UpdateProduct(
// 			tx,
// 			product,
// 			newProduct,
// 		)
// 		if err != nil {
// 			isSuccess = false
// 			logger.Logging(ctx, err)
// 		}
// 	}

// 	return err
// }

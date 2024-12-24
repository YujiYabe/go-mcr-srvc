package gateway

import (
	"backend/internal/4_domain/struct_object"
	"context"
)

// // GetVegetables ...
// func (receiver *Gateway) GetVegetables(ctx context.Context, requestVegetables map[string]int) error {
// 	err := receiver.ToPostgres.UpdateVegetables(ctx, requestVegetables)
// 	if err != nil {
// 		pkg.Logging(ctx, err)
// 		return err
// 	}

// 	return nil
// }

// // GetIngredients ...
// func (receiver *Gateway) GetIngredients(ctx context.Context, requestIngredients map[string]int) error {
// 	err := receiver.ToPostgres.UpdateIngredients(ctx, requestIngredients)
// 	if err != nil {
// 		pkg.Logging(ctx, err)
// 		return err
// 	}

// 	return nil
// }

// GetPersonList ...
func (receiver *Gateway) GetPersonList(
	ctx context.Context,

) (
	personList struct_object.PersonList,
	err error,
) {
	return receiver.ToPostgres.GetPersonList(ctx)
}

// GetPersonByCondition ...
func (receiver *Gateway) GetPersonByCondition(
	ctx context.Context,
	reqPerson struct_object.Person,
) (
	resPersonList struct_object.PersonList,
	err error,
) {
	return receiver.ToPostgres.GetPersonByCondition(
		ctx,
		reqPerson,
	)
}

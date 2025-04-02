package db_gateway

import (
	"context"

	"gorm.io/gorm"

	groupObject "backend/internal/4_domain/group_object"
)

type GatewayDB struct {
	ToPostgres ToPostgres
	ToRedis    ToRedis
}

// NewGatewayDB ...
func NewGatewayDB(
	toPostgres ToPostgres,
	toRedis ToRedis,
) *GatewayDB {
	return &GatewayDB{
		ToPostgres: toPostgres,
		ToRedis:    toRedis,
	}
}

type (

	// ToPostgres ...
	ToPostgres interface {
		WithOutTx(
			ctx context.Context,
		) (
			tx *gorm.DB,
		)

		BeginTx(
			ctx context.Context,
		) (
			tx *gorm.DB,
		)

		EndTx(
			ctx context.Context,
			tx *gorm.DB,
			isSuccess bool,
		) (
			err error,
		)

		GetPersonList(
			ctx context.Context,
			tx *gorm.DB,
		) (
			personList groupObject.PersonList,
		)

		GetPersonListByCondition(
			ctx context.Context,
			tx *gorm.DB,
			reqPerson groupObject.Person,
		) (
			resPersonList groupObject.PersonList,
		)
	}

	// ToRedis ...
	ToRedis interface {
		ResetPlaceListInRedis(
			ctx context.Context,
		) error
	}
)

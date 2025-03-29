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
		WithOutTx() (
			tx *gorm.DB,
		)

		BeginTx() (
			tx *gorm.DB,
		)

		EndTx(
			tx *gorm.DB,
			isSuccess bool,
		)

		GetPersonList(
			ctx context.Context,
		) (
			personList groupObject.PersonList,
		)

		GetPersonListByCondition(
			ctx context.Context,
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

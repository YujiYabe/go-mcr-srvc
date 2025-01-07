package usecase

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
	valueObject "backend/internal/4_domain/value_object"
)

// NewUseCase ...
func NewUseCase(
	toDomain ToDomain,
	toGateway ToGateway,
) ToUseCase {
	return &useCase{
		ToDomain:  toDomain,
		ToGateway: toGateway,
	}
}

type (
	// useCase ...
	useCase struct {
		ToDomain  ToDomain
		ToGateway ToGateway
	}

	// ToUseCase ...
	ToUseCase interface {
		Start()

		GetPersonList(
			ctx context.Context,
		) (
			personList groupObject.PersonList,
		)

		GetPersonByCondition(
			ctx context.Context,
			reqPerson groupObject.Person,
		) (
			resPersonList groupObject.PersonList,
		)

		FetchAccessToken(
			ctx context.Context,
			credential groupObject.Credential,
		) (
			accessToken valueObject.AccessToken,
		)

		ViaGRPC(
			ctx context.Context,
			reqPerson groupObject.Person,
		) (
			resPersonList groupObject.PersonList,
		)
	}

	// ToDomain ...
	ToDomain interface {
	}

	// ToGateway ...
	ToGateway interface {
		GetPersonList(
			ctx context.Context,
		) (
			personList groupObject.PersonList,
		)

		GetPersonByCondition(
			ctx context.Context,
			reqPerson groupObject.Person,
		) (
			resPersonList groupObject.PersonList,
		)

		FetchAccessToken(
			ctx context.Context,
			credential groupObject.Credential,
		) (
			accessToken valueObject.AccessToken,
		)

		ViaGRPC(
			ctx context.Context,
			reqPerson groupObject.Person,
		) (
			resPersonList groupObject.PersonList,
		)
	}
)

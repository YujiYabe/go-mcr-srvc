package usecase

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
	valueObject "backend/internal/4_domain/value_object"
)

// NewUseCase ...
func NewUseCase(
	toDomain ToDomain,
	toDBGateway ToDBGateway,
	toExternalGateway ToExternalGateway,
) ToUseCase {
	return &useCase{
		ToDomain:          toDomain,
		ToDBGateway:       toDBGateway,
		ToExternalGateway: toExternalGateway,
	}
}

type (
	// useCase ...
	useCase struct {
		ToDomain          ToDomain
		ToDBGateway       ToDBGateway
		ToExternalGateway ToExternalGateway
	}

	// ToUseCase ...
	ToUseCase interface {
		Start()

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

	// ToDBGateway ...
	ToDBGateway interface {
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

	// ToExternalGateway ...
	ToExternalGateway interface {
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

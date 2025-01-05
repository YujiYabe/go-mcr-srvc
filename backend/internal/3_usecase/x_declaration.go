package usecase

import (
	"context"

	domain "backend/internal/4_domain"
	structObject "backend/internal/4_domain/struct_object"
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
			personList structObject.PersonList,
			err error,
		)

		GetPersonByCondition(
			ctx context.Context,
			reqPerson structObject.Person,
		) (
			resPersonList structObject.PersonList,
			err error,
		)

		FetchAccessToken(
			ctx context.Context,
			credential structObject.Credential,
		) (
			accessToken valueObject.AccessToken,
			err error,
		)

		ViaGRPC(
			ctx context.Context,
		) (
			err error,
		)
	}

	// ToDomain ...
	ToDomain interface {
		ParseOrder(ctx context.Context, order *domain.Order) *domain.Assemble
		CookHamburgers(ctx context.Context, hamburgers []domain.Hamburger) error
	}

	// ToGateway ...
	ToGateway interface {
		GetPersonList(
			ctx context.Context,
		) (
			personList structObject.PersonList,
			err error,
		)

		GetPersonByCondition(
			ctx context.Context,
			reqPerson structObject.Person,
		) (
			resPersonList structObject.PersonList,
			err error,
		)

		FetchAccessToken(
			ctx context.Context,
			credential structObject.Credential,
		) (
			accessToken valueObject.AccessToken,
			err error,
		)

		ViaGRPC(
			ctx context.Context,
		) (
			err error,
		)
	}
)

package usecase

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
)

type (

	// ToUseCase ...
	ToUseCase interface {
		Start()

		GetPersonList(
			ctx context.Context,
		) (
			personList groupObject.PersonList,
			err error,
		)

		GetPersonListByCondition(
			ctx context.Context,
			reqPerson groupObject.Person,
		) (
			resPersonList groupObject.PersonList,
			err error,
		)

		FetchAccessToken(
			ctx context.Context,
			credential groupObject.Credential,
		) (
			accessToken typeObject.AccessToken,
			err error,
		)

		ViaGRPC(
			ctx context.Context,
			reqPerson groupObject.Person,
		) (
			resPersonList groupObject.PersonList,
			err error,
		)

		UpdatePerson(
			ctx context.Context,
			newPerson groupObject.Person,
		) error

		PublishTestTopic(
			ctx context.Context,
		) error
	}
)

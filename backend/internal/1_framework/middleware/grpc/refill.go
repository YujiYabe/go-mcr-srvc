package grpc_middleware

import (
	"context"
	"fmt"

	grpcParameter "backend/internal/1_framework/parameter/grpc"
	groupObject "backend/internal/4_domain/group_object"
)

func RefillUserGRPCToDomain(
	_ context.Context,
	request *grpcParameter.V1UserParameter,
) (
	reqUser *groupObject.User,
	err error,
) {
	var id *int
	if request.GetId() != 0 {
		id = new(int)
		*id = int(request.GetId())
	}

	var name *string
	if request.Name != nil {
		name = request.Name
	}

	var email *string
	if request.Email != nil {
		email = request.Email
	}

	reqUser, err = groupObject.NewUser(
		&groupObject.NewUserArgs{
			ID:    id,
			Name:  name,
			Email: email,
		},
	)

	return
}

func RefillUserDomainToGRPC(
	_ context.Context,
	userList groupObject.UserList,
) (
	v1UserParameterList []*grpcParameter.V1UserParameter,
	err error,
) {
	err = nil
	v1UserParameterList = []*grpcParameter.V1UserParameter{}

	for _, response := range userList.Content() {
		id32, convertErr := response.ID().ToUint32()
		if convertErr != nil {
			err = fmt.Errorf("convert user ID to uint32: %w", convertErr)

			return
		}
		name := response.Name().GetValue()
		email := response.Email().GetValue()
		v1UserParameter := &grpcParameter.V1UserParameter{
			Id:    &id32,
			Name:  &name,
			Email: &email,
		}
		v1UserParameterList = append(
			v1UserParameterList,
			v1UserParameter,
		)
	}

	return
}

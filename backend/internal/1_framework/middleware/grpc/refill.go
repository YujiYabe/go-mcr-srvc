package grpc_middleware

import (
	"context"

	grpcParameter "backend/internal/1_framework/parameter/grpc"
	groupObject "backend/internal/4_domain/group_object"
)

func RefillPersonGRPCToDomain(
	ctx context.Context,
	request *grpcParameter.V1PersonParameter,
) (
	reqPerson *groupObject.Person,
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

	var mailAddress *string
	if request.MailAddress != nil {
		mailAddress = request.MailAddress
	}

	reqPerson = groupObject.NewPerson(
		ctx,
		&groupObject.NewPersonArgs{
			ID:          id,
			Name:        name,
			MailAddress: mailAddress,
		},
	)

	return
}

func RefillPersonDomainToGRPC(
	ctx context.Context,
	personList groupObject.PersonList,
) (
	v1PersonParameterList []*grpcParameter.V1PersonParameter,

) {
	v1PersonParameterList = []*grpcParameter.V1PersonParameter{}

	for _, response := range personList.Content {
		id32 := uint32(response.ID.GetValue())
		name := response.Name.GetValue()
		mailAddress := response.MailAddress.GetValue()
		v1PersonParameter := &grpcParameter.V1PersonParameter{
			Id:          &id32,
			Name:        &name,
			MailAddress: &mailAddress,
		}
		v1PersonParameterList = append(
			v1PersonParameterList,
			v1PersonParameter,
		)
	}

	return
}

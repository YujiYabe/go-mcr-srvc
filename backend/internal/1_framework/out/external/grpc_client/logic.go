package grpc_client

import (
	"context"

	grpcMiddleware "backend/internal/1_framework/middleware/grpc"
	grpcParameter "backend/internal/1_framework/parameter/grpc"
	groupObject "backend/internal/4_domain/group_object"
)

// ...
// ViaGRPC ...
func (receiver *GRPCClient) ViaGRPC(
	ctx context.Context,
	reqUser groupObject.User,
) (
	resUserList groupObject.UserList,
	err error,
) {
	// traceID := requestContextMiddleware.GetRequestContext(ctx).TraceID.GetValue()
	// logger.Logging(ctx, traceID)

	// クライアントの作成
	client := grpcParameter.NewUserServiceClient(receiver.Conn)

	// リクエストの作成
	v1GetUserByConditionRequest := &grpcParameter.GetUserListByConditionRequest{
		V1UserParameter: &grpcParameter.V1UserParameter{},
	}

	if !reqUser.Name().GetIsNil() && reqUser.Name().GetValue() != "" {
		value := reqUser.Name().GetValue()
		v1GetUserByConditionRequest.V1UserParameter.Name = &value
	}

	if !reqUser.Email().GetIsNil() && reqUser.Email().GetValue() != "" {
		value := reqUser.Email().GetValue()
		v1GetUserByConditionRequest.V1UserParameter.Email = &value
	}

	ctx = grpcMiddleware.ContextToMetadata(ctx)

	// gRPCリクエストの実行
	grpcUserList, err := client.GetUserListByCondition(
		ctx,
		v1GetUserByConditionRequest,
	)
	if err != nil {
		return
	}
	userArgs := make([]groupObject.NewUserArgs, 0, len(grpcUserList.V1UserParameterArray.Users))
	for _, grpcUser := range grpcUserList.V1UserParameterArray.Users {
		id := int(grpcUser.GetId())
		name := grpcUser.GetName()
		email := grpcUser.GetEmail()
		userArgs = append(userArgs, groupObject.NewUserArgs{
			ID:    &id,
			Name:  &name,
			Email: &email,
		})
	}

	// traceID = requestContextMiddleware.GetRequestContext(ctx).TraceID.GetValue()
	// logger.Logging(ctx, traceID)

	return groupObject.NewUserList(&groupObject.NewUserListArgs{
		Content: userArgs,
	})

}

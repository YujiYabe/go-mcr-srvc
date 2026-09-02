package goGRPC

import (
	"context"
	"time"

	grpcMiddleware "backend/internal/1_framework/middleware/grpc"
	middlewareRequestContext "backend/internal/1_framework/middleware/request_context"
	grpcParameter "backend/internal/1_framework/parameter/grpc"
	"backend/internal/logger"
)

// GoGRPC ...
type GoGRPC struct {
	Server
	address string
}

// ------------
func (receiver *Server) GetUserListByCondition(
	ctx context.Context,
	getUserListByConditionRequest *grpcParameter.GetUserListByConditionRequest,
) (
	v1GetUserListByConditionResponse *grpcParameter.GetUserListByConditionResponse,
	err error,
) {
	v1GetUserListByConditionResponse = nil
	err = nil
	requestContext := middlewareRequestContext.GetRequestContext(ctx)
	if requestContext == nil {
		v1GetUserListByConditionResponse, err = nil, ctx.Err()
		return
	}

	timeoutMillSecond := requestContext.TimeOutMillSecond().GetValue()

	ctx, cancel := context.WithTimeout(
		ctx,
		time.Duration(timeoutMillSecond)*time.Millisecond,
	)
	defer cancel() // コンテキストのキャンセルを必ず呼び出す

	done := make(chan struct{})

	// ゴルーチンで処理を実行
	go func() {
		v1GetUserListByConditionResponse, err = receiver.getUserListByCondition(
			ctx,
			getUserListByConditionRequest,
		)
		close(done) // 処理完了を通知
	}()

	// タイムアウトまたは処理完了を待つ
	select {
	case <-done:
		// 処理が完了した場合
		return

	case <-ctx.Done():
		// タイムアウトした場合
		logger.Logging(ctx, ctx.Err())
		v1GetUserListByConditionResponse, err = nil, ctx.Err()
		return
	}
}

func (receiver *Server) getUserListByCondition(
	ctx context.Context,
	getUserListByConditionRequest *grpcParameter.GetUserListByConditionRequest,
) (
	getUserListByConditionResponse *grpcParameter.GetUserListByConditionResponse,
	err error,
) {
	getUserListByConditionResponse = &grpcParameter.GetUserListByConditionResponse{}

	reqUser, err := grpcMiddleware.RefillUserGRPCToDomain(
		ctx,
		getUserListByConditionRequest.GetV1UserParameter(),
	)
	if err != nil {
		logger.Logging(ctx, err)
		getUserListByConditionResponse = nil
		return
	}

	responseList, err := receiver.Controller.GetUserListByCondition(
		ctx,
		*reqUser,
	)
	if err != nil {
		logger.Logging(ctx, err)
		getUserListByConditionResponse = nil
		return
	}

	v1UserParameterArray := &grpcParameter.V1UserParameterArray{}
	v1UserParameterArray.Users, err = grpcMiddleware.RefillUserDomainToGRPC(
		ctx,
		responseList,
	)
	if err != nil {
		logger.Logging(ctx, err)
		getUserListByConditionResponse = nil
		return
	}

	getUserListByConditionResponse.V1UserParameterArray = v1UserParameterArray

	// logger.Logging(ctx, traceID)

	return
}

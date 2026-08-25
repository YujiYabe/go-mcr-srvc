package goGRPC

import (
	"context"
	"time"

	grpcMiddleware "backend/internal/1_framework/middleware/grpc"
	requestContextMiddleware "backend/internal/1_framework/middleware/request_context"
	grpcParameter "backend/internal/1_framework/parameter/grpc"
	"backend/internal/logger"
)

// GoGRPC ...
type GoGRPC struct {
	Server
}

// ------------
func (receiver *Server) GetPersonListByCondition(
	ctx context.Context,
	getPersonListByConditionRequest *grpcParameter.GetPersonListByConditionRequest,
) (
	v1GetPersonListByConditionResponse *grpcParameter.GetPersonListByConditionResponse,
	err error,
) {
	requestContext := requestContextMiddleware.GetRequestContext(ctx)
	if requestContext == nil {
		return nil, ctx.Err()
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
		v1GetPersonListByConditionResponse, err = receiver.getPersonListByCondition(
			ctx,
			getPersonListByConditionRequest,
		)
		close(done) // 処理完了を通知
	}()

	// タイムアウトまたは処理完了を待つ
	select {
	case <-done:
		// 処理が完了した場合
		return v1GetPersonListByConditionResponse, err

	case <-ctx.Done():
		// タイムアウトした場合
		logger.Logging(ctx, ctx.Err())
		return nil, ctx.Err()
	}
}

func (receiver *Server) getPersonListByCondition(
	ctx context.Context,
	getPersonListByConditionRequest *grpcParameter.GetPersonListByConditionRequest,
) (
	getPersonListByConditionResponse *grpcParameter.GetPersonListByConditionResponse,
	err error,
) {
	getPersonListByConditionResponse = &grpcParameter.GetPersonListByConditionResponse{}

	// traceID := requestContextMiddleware.GetRequestContext(ctx).TraceID.GetValue()
	// logger.Logging(ctx, traceID)

	reqPerson, err := grpcMiddleware.RefillPersonGRPCToDomain(
		ctx,
		getPersonListByConditionRequest.GetV1PersonParameter(),
	)
	if err != nil {
		logger.Logging(ctx, err)
		return nil, err
	}

	responseList, err := receiver.Controller.GetPersonListByCondition(
		ctx,
		*reqPerson,
	)
	if err != nil {
		logger.Logging(ctx, err)
		return nil, err
	}

	v1PersonParameterArray := &grpcParameter.V1PersonParameterArray{}
	v1PersonParameterArray.Persons = grpcMiddleware.RefillPersonDomainToGRPC(
		ctx,
		responseList,
	)

	getPersonListByConditionResponse.V1PersonParameterArray = v1PersonParameterArray

	// logger.Logging(ctx, traceID)

	return
}

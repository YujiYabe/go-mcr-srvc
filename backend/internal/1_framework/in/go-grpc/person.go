package goGRPC

import (
	"context"
	"fmt"
	"time"

	grpcMiddleware "backend/internal/1_framework/middleware/grpc"
	grpcParameter "backend/internal/1_framework/parameter/grpc"
	valueObject "backend/internal/4_domain/value_object"

	"backend/pkg"
)

// GoGRPC ...
type GoGRPC struct {
	Server
}

// ------------
func (receiver *Server) GetPersonByCondition(
	ctx context.Context,
	getPersonByConditionRequest *grpcParameter.GetPersonByConditionRequest,
) (
	v1GetPersonByConditionResponse *grpcParameter.GetPersonByConditionResponse,
	err error,
) {
	ctx = grpcMiddleware.CommonToContext(
		ctx,
		getPersonByConditionRequest.GetV1CommonParameter(),
	)

	timeoutSecond := valueObject.GetTimeoutSecond(ctx)
	ctx, cancel := context.WithTimeout(
		ctx,
		time.Duration(timeoutSecond)*time.Millisecond,
	)
	defer cancel() // コンテキストのキャンセルを必ず呼び出す

	done := make(chan struct{})

	// ゴルーチンで処理を実行
	go func() {
		v1GetPersonByConditionResponse, err = receiver.processPersonRequest(
			ctx,
			getPersonByConditionRequest,
		)
		close(done) // 処理完了を通知
	}()

	// タイムアウトまたは処理完了を待つ
	select {
	case <-done:
		// 処理が完了した場合
		fmt.Println("正常に終了しました")
		return v1GetPersonByConditionResponse, err

	case <-ctx.Done():
		// タイムアウトした場合
		fmt.Println("タイムアウトしました")
		return nil, ctx.Err()
	}
}

func (receiver *Server) processPersonRequest(
	ctx context.Context,
	getPersonByConditionRequest *grpcParameter.GetPersonByConditionRequest,
) (
	v1GetPersonByConditionResponse *grpcParameter.GetPersonByConditionResponse,
	err error,
) {
	v1GetPersonByConditionResponse = &grpcParameter.GetPersonByConditionResponse{}

	traceID := valueObject.GetTraceID(ctx)
	pkg.Logging(ctx, traceID)

	reqPerson := grpcMiddleware.RefillPersonGRPCToDomain(
		ctx,
		getPersonByConditionRequest.GetV1PersonParameter(),
	)
	if reqPerson.GetError() != nil {
		pkg.Logging(ctx, reqPerson.GetError())
		return nil, reqPerson.GetError()
	}

	responseList := receiver.Controller.GetPersonByCondition(
		ctx,
		*reqPerson,
	)
	if responseList.GetError() != nil {
		pkg.Logging(ctx, responseList.GetError())
		return nil, responseList.GetError()
	}

	v1PersonParameterArray := &grpcParameter.V1PersonParameterArray{}
	v1PersonParameterArray.Persons = grpcMiddleware.RefillPersonDomainToGRPC(
		ctx,
		responseList,
	)

	v1GetPersonByConditionResponse.V1PersonParameterArray = v1PersonParameterArray
	v1GetPersonByConditionResponse.V1CommonParameter = &grpcParameter.V1CommonParameter{
		Immutable: &grpcParameter.V1ImmutableParameter{
			TraceId: traceID,
		},
		Mutable: &grpcParameter.V1MutableParameter{
			TimeStamp: time.Now().Format(time.RFC3339),
		},
	}

	pkg.Logging(ctx, traceID)

	return v1GetPersonByConditionResponse, nil
}

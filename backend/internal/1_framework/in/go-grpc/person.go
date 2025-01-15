package goGRPC

import (
	"context"
	"time"

	grpcMiddleware "backend/internal/1_framework/middleware/grpc"
	grpcParameter "backend/internal/1_framework/parameter/grpc"
	groupObject "backend/internal/4_domain/group_object"

	"backend/pkg"
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
	requestContext := groupObject.GetRequestContext(ctx)
	if requestContext.GetError() != nil {
		return nil, requestContext.GetError()
	}

	timeoutSecond := requestContext.TimeOutSecond.GetValue()
	// requestStartTime := requestContext.RequestStartTime.GetValue()
	// log.Println("== == == == == == == == == == ")
	// pkg.Logging(ctx, timeoutSecond)
	// pkg.Logging(ctx, requestStartTime)
	// log.Println("== == == == == == == == == == ")
	// now := time.Now().UnixMilli()

	// pkg.Logging(ctx, "-- -- -- -- -- -- -- -- -- -- ")
	// pkg.Logging(ctx, time.UnixMilli(requestStartTime).Format("2006-01-02 15:04:05.000"))
	// pkg.Logging(ctx, time.UnixMilli(now).Format("2006-01-02 15:04:05.000"))

	ctx, cancel := context.WithTimeout(
		ctx,
		time.Duration(timeoutSecond)*time.Millisecond,
	)
	defer cancel() // コンテキストのキャンセルを必ず呼び出す

	done := make(chan struct{})

	// ゴルーチンで処理を実行
	go func() {
		v1GetPersonListByConditionResponse, err = receiver.processPersonRequest(
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
		pkg.Logging(ctx, ctx.Err())
		return nil, ctx.Err()
	}
}

func (receiver *Server) processPersonRequest(
	ctx context.Context,
	getPersonListByConditionRequest *grpcParameter.GetPersonListByConditionRequest,
) (
	v1GetPersonListByConditionResponse *grpcParameter.GetPersonListByConditionResponse,
	err error,
) {
	v1GetPersonListByConditionResponse = &grpcParameter.GetPersonListByConditionResponse{}

	// traceID := groupObject.GetRequestContext(ctx).TraceID.GetValue()
	// pkg.Logging(ctx, traceID)

	reqPerson := grpcMiddleware.RefillPersonGRPCToDomain(
		ctx,
		getPersonListByConditionRequest.GetV1PersonParameter(),
	)
	if reqPerson.GetError() != nil {
		pkg.Logging(ctx, reqPerson.GetError())
		return nil, reqPerson.GetError()
	}

	responseList := receiver.Controller.GetPersonListByCondition(
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

	v1GetPersonListByConditionResponse.V1PersonParameterArray = v1PersonParameterArray

	// pkg.Logging(ctx, traceID)

	return v1GetPersonListByConditionResponse, nil
}

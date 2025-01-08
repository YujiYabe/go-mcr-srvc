package grpc_middleware

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"

	grpcParameter "backend/internal/1_framework/parameter/grpc"
	valueObject "backend/internal/4_domain/value_object"
	"backend/pkg"
)

func CommonToContext(
	ctx context.Context,
	req *grpcParameter.V1CommonParameter,
) context.Context {

	// エラーがあれば都度処理
	if req.GetError() != nil {
		log.Println("== == == == == == == == == == ")
		pkg.Logging(ctx, req.GetError())
		log.Println("== == == == == == == == == == ")
	}

	// 不変データがなければ追加
	if req.GetImmutable() != nil {
		ctx = traceIDToContext(ctx, req.GetImmutable())
		ctx = requestStartTimeToContext(ctx, req.GetImmutable())

	}

	//  可変データの更新または追加
	ctx = timeStampToContext(ctx, req.GetMutable())
	// ctx = TimeoutSecondToContext(ctx, req.GetImmutable())

	log.Println("== == == == == == == == == == ")
	pkg.Logging(ctx, req.GetMutable())
	log.Println("== == == == == == == == == == ")

	return ctx
}

func timeStampToContext(
	ctx context.Context,
	v1IMutableParameter *grpcParameter.V1MutableParameter,
) (
	newCtx context.Context,
) {
	timesStamp := v1IMutableParameter.GetTimeStamp()

	// traceID をコンテキストに追加
	newCtx = context.WithValue(
		ctx,
		valueObject.TimeStampContextName,
		timesStamp,
	)

	return
}

func requestStartTimeToContext(
	ctx context.Context,
	v1ImmutableParameter *grpcParameter.V1ImmutableParameter,
) (
	newCtx context.Context,
) {
	requestStartTime := v1ImmutableParameter.GetRequestStartTime()

	// requestStartTime が無い場合は新規生成
	if requestStartTime == 0 {
		requestStartTime = time.Now().Unix()
	}

	// requestStartTime をコンテキストに追加
	newCtx = context.WithValue(
		ctx,
		valueObject.RequestStartTimeContextName,
		requestStartTime,
	)

	return
}

func traceIDToContext(
	ctx context.Context,
	v1ImmutableParameter *grpcParameter.V1ImmutableParameter,
) (
	newCtx context.Context,
) {
	traceID := v1ImmutableParameter.GetTraceId()

	// traceID が無い場合は新規生成
	if traceID == "" {
		traceID = uuid.New().String()
	}

	// traceID をコンテキストに追加
	newCtx = context.WithValue(
		ctx,
		valueObject.TraceIDContextName,
		traceID,
	)

	return
}

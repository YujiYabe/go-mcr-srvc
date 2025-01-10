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

// ------------
func CommonToContext(
	ctx context.Context,
	req *grpcParameter.V1CommonParameter,
) context.Context {

	// エラーがあれば都度処理
	if req.GetV1Error() != nil {
		log.Println("== == == == == == == == == == ")
		pkg.Logging(ctx, req.GetV1Error())
		log.Println("== == == == == == == == == == ")
	}

	// 不変データがなければ追加
	// ctx = traceIDToContext(ctx, req.GetImmutable())
	// ctx = requestStartTimeToContext(ctx, req.GetImmutable())
	ctx = traceIDToContext(ctx, req.GetV1RequestContext())
	ctx = requestStartTimeToContext(ctx, req.GetV1RequestContext())

	//  可変データの更新または追加
	ctx = timeoutSecondToContext(ctx)

	return ctx
}

// リクエスト処理の残り時間（秒）を計算
func timeoutSecondToContext(
	ctx context.Context,
) (
	newCtx context.Context,
) {
	requestStartTime := valueObject.GetRequestStartTime(ctx)
	currentTimestamp := time.Now().UnixMilli()
	requestEndTime := time.UnixMilli(requestStartTime).Add(5 * time.Second).UnixMilli()
	timeoutSecond := requestEndTime - currentTimestamp

	// requestStartFormatted := time.UnixMilli(requestStartTime).Format("20060102-150405.000")
	// currentFormatted := time.UnixMilli(currentTimestamp).Format("20060102-150405.000")
	// endTimeFormatted := time.UnixMilli(requestEndTime).Format("20060102-150405.000")

	// log.Println("== == == == == == == == == == ")
	// log.Printf("Request Start Time: %s\n", requestStartFormatted)
	// log.Println("== == == == == == == == == == ")
	// log.Printf("Current Time: %s\n", currentFormatted)
	// log.Println("== == == == == == == == == == ")
	// log.Printf("Request End Time: %s\n", endTimeFormatted)
	// log.Println("== == == == == == == == == == ")

	// requestStartTime をコンテキストに追加
	newCtx = context.WithValue(
		ctx,
		valueObject.TimeOutSecondContextName,
		timeoutSecond,
	)

	return
}

// ------------
func requestStartTimeToContext(
	ctx context.Context,
	v1RequestContext *grpcParameter.V1RequestContext,
) (
	newCtx context.Context,
) {
	requestStartTime := v1RequestContext.GetRequestStartTime()

	// requestStartTime が無い場合は新規生成
	if requestStartTime == 0 {
		requestStartTime = time.Now().UnixMilli()
	}

	// requestStartTime をコンテキストに追加
	newCtx = context.WithValue(
		ctx,
		valueObject.RequestStartTimeContextName,
		requestStartTime,
	)

	return
}

// ------------
func traceIDToContext(
	ctx context.Context,
	v1RequestContext *grpcParameter.V1RequestContext,
) (
	newCtx context.Context,
) {
	traceID := v1RequestContext.GetTraceId()

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

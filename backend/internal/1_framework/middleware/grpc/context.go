package grpc_middleware

import (
	"context"

	"github.com/google/uuid"

	grpcParameter "backend/internal/1_framework/parameter/grpc"
	valueObject "backend/internal/4_domain/value_object"
)

func CommonToContext(
	ctx context.Context,
	v1CommonParameter *grpcParameter.V1CommonParameter,
) context.Context {

	ctx = traceIDToContext(ctx, v1CommonParameter)
	// ctx = requestStartTimeToContext(ctx)
	// ctx = timestampToContext(ctx)

	return ctx
}

func traceIDToContext(
	ctx context.Context,
	v1CommonParameter *grpcParameter.V1CommonParameter,
) (
	newCtx context.Context,
) {
	traceID := v1CommonParameter.Immutable.GetTraceId()

	// リクエストIDが無い場合は新規生成
	if traceID == "" {
		traceID = uuid.New().String()
	}

	// リクエストIDをコンテキストに追加
	newCtx = context.WithValue(
		ctx,
		valueObject.TraceIDContextName,
		traceID,
	)

	return
}

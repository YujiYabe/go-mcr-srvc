package grpc_middleware

import (
	"context"
	"log"

	"github.com/google/uuid"

	grpcParameter "backend/internal/1_framework/parameter/grpc"
	valueObject "backend/internal/4_domain/value_object"
	"backend/pkg"
)

func CommonToContext(
	ctx context.Context,
	req *grpcParameter.GetPersonByConditionRequest,
) context.Context {

	// エラーがあれば都度処理
	if req.V1CommonParameter.GetError() != nil {
		log.Println("== == == == == == == == == == ")
		pkg.Logging(ctx, req.V1CommonParameter.GetError())
		log.Println("== == == == == == == == == == ")
	}

	// 不変データがなければ追加
	if req.V1CommonParameter.GetImmutable() != nil {
		ctx = traceIDToContext(ctx, req.V1CommonParameter.GetImmutable())

		log.Println("== == == == == == == == == == ")
		pkg.Logging(ctx, req.V1CommonParameter.GetImmutable())
		log.Println("== == == == == == == == == == ")
	}

	//  可変データの更新または追加
	if req.V1CommonParameter.GetMutable() != nil {
		ctx = traceIDToContext(ctx, req.V1CommonParameter.GetImmutable())

		log.Println("== == == == == == == == == == ")
		pkg.Logging(ctx, req.V1CommonParameter.GetMutable())
		log.Println("== == == == == == == == == == ")

	}

	return ctx
}

func traceIDToContext(
	ctx context.Context,
	v1ImmutableParameter *grpcParameter.V1ImmutableParameter,
) (
	newCtx context.Context,
) {
	traceID := v1ImmutableParameter.GetTraceId()

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

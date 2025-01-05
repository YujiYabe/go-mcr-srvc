package grpc_middleware

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"

	valueObject "backend/internal/4_domain/value_object"
)

func timestampToContext(
	nowCtx context.Context,
	md metadata.MD,
) (
	newCtx context.Context,
) {
	var traceID string

	values := md.Get(string(valueObject.TimeStampMetaName))
	if len(values) > 0 {
		traceID = values[0]
	}

	// リクエストIDが無い場合は新規生成
	if traceID == "" {
		traceID = uuid.New().String()
	}

	// リクエストIDをコンテキストに追加
	newCtx = context.WithValue(
		nowCtx,
		valueObject.TimeStampContextName,
		traceID,
	)

	return
}

package grpc_middleware

import (
	"context"
	"strconv"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"

	valueObject "backend/internal/4_domain/value_object"
)

func traceIDToContext(
	nowCtx context.Context,
	md metadata.MD,
) (
	newCtx context.Context,
) {
	var traceID string

	values := md.Get(string(valueObject.TraceIDMetaName))
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
		valueObject.TraceIDContextName,
		traceID,
	)

	return
}

func requestStartTimeToContext(
	nowCtx context.Context,
	md metadata.MD,
) (
	newCtx context.Context,
) {
	var requestStartTime int64

	values := md.Get(string(valueObject.RequestStartTimeMetaName))
	if len(values) > 0 {
		parsedTime, err := strconv.ParseInt(values[0], 10, 64)
		if err == nil {
			requestStartTime = parsedTime
		}
	}

	// リクエスト開始時間が無い場合は現在時刻を使用
	if requestStartTime == 0 {
		requestStartTime = time.Now().UnixNano()
	}
	// リクエスト開始時間をコンテキストに追加
	newCtx = context.WithValue(
		nowCtx,
		valueObject.RequestStartTimeContextName,
		requestStartTime,
	)

	return
}

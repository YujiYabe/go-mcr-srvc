package pkg

import (
	"context"
)

type contextKey string

const RequestIDKey contextKey = "requestID"

func GetNewContext(
	ctx context.Context,
	headerXRequestID string,
) (
	newCtx context.Context,
) {
	// echoのX-Request-IDをcontextに埋め込み
	newCtx = context.WithValue(
		ctx,              // parent
		RequestIDKey,     // key
		headerXRequestID, // val
	)

	return
}

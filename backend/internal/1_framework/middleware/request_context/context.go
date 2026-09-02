package request_context

import (
	"context"
)

func GetRequestContext(
	ctx context.Context,
) (
	value *RequestContext,
) {
	value = nil
	requestContext, ok := ctx.Value(RequestContextContextName).(RequestContext)
	if ok {
		value = &requestContext
	}

	return
}

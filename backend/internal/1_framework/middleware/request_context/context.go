package request_context

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
)

func GetRequestContext(
	ctx context.Context,
) (
	value *groupObject.RequestContext,
) {
	requestContext, ok := ctx.Value(groupObject.RequestContextContextName).(groupObject.RequestContext)
	if ok {
		value = &requestContext
	}

	return
}

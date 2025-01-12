package grpc_middleware

import (
	"context"
	"fmt"
	"log"

	grpcParameter "backend/internal/1_framework/parameter/grpc"
	groupObject "backend/internal/4_domain/group_object"
	"backend/pkg"
)

// ------------
func CommonToContext(
	ctx context.Context,
	req *grpcParameter.V1CommonParameter,
) context.Context {

	// エラーがあれば都度処理
	if req.GetV1Error() != nil {
		pkg.Logging(ctx, req.GetV1Error())
	}

	newRequestContextArgs := &groupObject.NewRequestContextArgs{
		RequestStartTime: &req.V1RequestContext.RequestStartTime,
		TraceID:          &req.V1RequestContext.TraceId,
		ClientIP:         &req.V1RequestContext.ClientIp,
		UserAgent:        &req.V1RequestContext.UserAgent,
		UserID:           &req.V1RequestContext.UserId,
		AccessToken:      &req.V1RequestContext.AccessToken,
		TenantID:         &req.V1RequestContext.TenantId,
		Locale:           &req.V1RequestContext.Locale,
		Timezone:         &req.V1RequestContext.Timezone,
	}

	requestContext := groupObject.NewRequestContext(
		ctx,
		newRequestContextArgs,
	)
	debug := requestContext
	fmt.Println(" ----------------------------------- ")
	fmt.Printf("%+v\n", debug.TraceID.GetValue())
	fmt.Println(" ----------------------------------- ")

	if requestContext.GetError() != nil {
		log.Println(requestContext.GetError())
		return ctx
	}
	ctx = context.WithValue(
		ctx,
		groupObject.RequestContextContextName,
		*requestContext,
	)

	return ctx
}

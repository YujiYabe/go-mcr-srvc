package grpc_middleware

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"

	grpcParameter "backend/internal/1_framework/parameter/grpc"
	groupObject "backend/internal/4_domain/group_object"
	valueObject "backend/internal/4_domain/value_object"
	"backend/pkg"
)

// func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
// 	return func(
// 		ctx context.Context,
// 		req interface{},
// 		info *grpc.UnaryServerInfo,
// 		handler grpc.UnaryHandler,
// 	) (
// 		interface{},
// 		error,
// 	) {
// 		if commonReq, ok := req.(*grpcParameter.V1CommonParameter); ok {
// 			ctx = CommonToContext(ctx, commonReq)
// 			log.Println("== == == == == == == == == == ")
// 			pkg.Logging(ctx, "ok")
// 			log.Println("== == == == == == == == == == ")
// 		} else {
// 			log.Println("== == == == == == == == == == ")
// 			pkg.Logging(ctx, "ng")
// 			log.Println("== == == == == == == == == == ")
// 		}

// 		return handler(ctx, req)
// 	}
// }

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
		TimeZone:         &req.V1RequestContext.TimeZone,
	}

	requestContext := groupObject.NewRequestContext(
		ctx,
		newRequestContextArgs,
	)

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

// ------------
func ContextToMetadata(
	ctx context.Context,
) context.Context {

	metaDataMap := map[string]string{}

	metaDataMap[string(valueObject.TraceIDMetaName)] = groupObject.GetRequestContext(ctx).TraceID.GetValue()
	metaDataMap[string(valueObject.ClientIPMetaName)] = groupObject.GetRequestContext(ctx).ClientIP.GetValue()
	metaDataMap[string(valueObject.UserAgentMetaName)] = groupObject.GetRequestContext(ctx).UserAgent.GetValue()
	metaDataMap[string(valueObject.UserIDMetaName)] = groupObject.GetRequestContext(ctx).UserID.GetValue()
	metaDataMap[string(valueObject.AccessTokenMetaName)] = groupObject.GetRequestContext(ctx).AccessToken.GetValue()
	metaDataMap[string(valueObject.TenantIDMetaName)] = groupObject.GetRequestContext(ctx).TenantID.GetValue()
	metaDataMap[string(valueObject.LocaleMetaName)] = groupObject.GetRequestContext(ctx).Locale.GetValue()
	metaDataMap[string(valueObject.TimeZoneMetaName)] = groupObject.GetRequestContext(ctx).TimeZone.GetValue()

	md := metadata.New(
		metaDataMap,
	)

	ctx = metadata.NewOutgoingContext(ctx, md)

	return ctx
}

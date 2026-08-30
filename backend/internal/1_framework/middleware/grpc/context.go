package grpc_middleware

import (
	"context"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	middlewareRequestContext "backend/internal/1_framework/middleware/request_context"
	typeObject "backend/internal/4_domain/type_object"
	"backend/internal/logger"
)

// ------------
func ContextToMetadata(
	ctx context.Context,
) (
	ctxResult context.Context,
) {
	requestContext := middlewareRequestContext.GetRequestContext(ctx)
	if requestContext == nil {
		return ctx
	}

	metaDataMap := map[string]string{}

	// string value
	metaDataMap[string(typeObject.TraceIDHeaderName)] = requestContext.TraceID().GetValue()
	metaDataMap[string(typeObject.ClientIPHeaderName)] = requestContext.ClientIP().GetValue()
	metaDataMap[string(typeObject.UserAgentHeaderName)] = requestContext.UserAgent().GetValue()
	metaDataMap[string(typeObject.UserIDHeaderName)] = requestContext.UserID().GetValue()
	metaDataMap[string(typeObject.AccessTokenHeaderName)] = requestContext.AccessToken().GetValue()
	metaDataMap[string(typeObject.TenantIDHeaderName)] = requestContext.TenantID().GetValue()
	metaDataMap[string(typeObject.LocaleHeaderName)] = requestContext.Locale().GetValue()
	metaDataMap[string(typeObject.TimeZoneHeaderName)] = requestContext.TimeZone().GetValue()

	// int64 value
	metaDataMap[string(typeObject.RequestStartTimeHeaderName)] = requestContext.RequestStartTime().GetString()

	// permissionListを文字列のスライスとして格納
	metaDataMap[string(typeObject.PermissionListHeaderName)] = strings.Join(
		requestContext.PermissionList().GetSliceValue(),
		",",
	)

	metadataCollection := metadata.New(
		metaDataMap,
	)

	ctx = metadata.NewOutgoingContext(ctx, metadataCollection)

	return ctx
}

func MetadataToContext(
	ctx context.Context,
) (
	ctxResult context.Context,
) {
	metadataCollection, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	newRequestContextArgs := &middlewareRequestContext.NewRequestContextArgs{}

	// ________________________________
	if len(metadataCollection.Get(string(typeObject.RequestStartTimeHeaderName))) != 0 {
		requestStartTime, _ := strconv.ParseInt(
			metadataCollection.Get(string(typeObject.RequestStartTimeHeaderName))[0], 10, 64,
		)
		newRequestContextArgs.RequestStartTime = &requestStartTime
	}

	// ________________________________
	if len(metadataCollection.Get(string(typeObject.PermissionListHeaderName))) != 0 {
		permissionList := []string{}
		for _, permissionHeader := range metadataCollection.Get(string(typeObject.PermissionListHeaderName)) {
			for _, permission := range strings.Split(permissionHeader, ",") {
				if permission == "" {
					continue
				}
				permissionList = append(permissionList, permission)
			}
		}
		newRequestContextArgs.PermissionList = permissionList

	}

	// ________________________________
	if len(metadataCollection.Get(string(typeObject.TraceIDHeaderName))) != 0 {
		value := metadataCollection.Get(string(typeObject.TraceIDHeaderName))[0]
		newRequestContextArgs.TraceID = &value
	}

	// ________________________________
	if len(metadataCollection.Get(string(typeObject.ClientIPHeaderName))) != 0 {
		value := metadataCollection.Get(string(typeObject.ClientIPHeaderName))[0]
		newRequestContextArgs.ClientIP = &value
	}

	// ________________________________
	if len(metadataCollection.Get(string(typeObject.UserAgentHeaderName))) != 0 {
		value := metadataCollection.Get(string(typeObject.UserAgentHeaderName))[0]
		newRequestContextArgs.UserAgent = &value
	}

	// ________________________________
	if len(metadataCollection.Get(string(typeObject.UserIDHeaderName))) != 0 {
		value := metadataCollection.Get(string(typeObject.UserIDHeaderName))[0]
		newRequestContextArgs.UserID = &value
	}

	// ________________________________
	if len(metadataCollection.Get(string(typeObject.AccessTokenHeaderName))) != 0 {
		value := metadataCollection.Get(string(typeObject.AccessTokenHeaderName))[0]
		newRequestContextArgs.AccessToken = &value
	}

	// ________________________________
	if len(metadataCollection.Get(string(typeObject.TenantIDHeaderName))) != 0 {
		value := metadataCollection.Get(string(typeObject.TenantIDHeaderName))[0]
		newRequestContextArgs.TenantID = &value
	}

	// ________________________________
	if len(metadataCollection.Get(string(typeObject.LocaleHeaderName))) != 0 {
		value := metadataCollection.Get(string(typeObject.LocaleHeaderName))[0]
		newRequestContextArgs.Locale = &value
	}

	// ________________________________
	if len(metadataCollection.Get(string(typeObject.TimeZoneHeaderName))) != 0 {
		value := metadataCollection.Get(string(typeObject.TimeZoneHeaderName))[0]
		newRequestContextArgs.TimeZone = &value
	}

	requestContext, err := middlewareRequestContext.NewRequestContext(
		newRequestContextArgs,
	)
	if err != nil {
		logger.Logging(ctx, err)
		return ctx
	}

	ctx = context.WithValue(
		ctx,
		middlewareRequestContext.RequestContextContextName,
		*requestContext,
	)

	// ________________________________
	// logで追跡するために、contextにTraceIDを設定する
	ctx = context.WithValue(
		ctx,
		typeObject.TraceIDContextName,
		requestContext.TraceID().GetValue(),
	)

	return ctx
}

func UnaryServerInterceptor() (
	unaryServerInterceptor grpc.UnaryServerInterceptor,
) {
	return func(
		ctx context.Context,
		req interface{},
		_ *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (
		interface{},
		error,
	) {
		ctx = MetadataToContext(ctx)

		return handler(ctx, req)
	}
}

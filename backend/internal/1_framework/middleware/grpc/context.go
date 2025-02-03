package grpc_middleware

import (
	"context"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	groupObject "backend/internal/4_domain/group_object"
	valueObject "backend/internal/4_domain/value_object"
	"backend/internal/logger"
)

// ------------
func ContextToMetadata(
	ctx context.Context,
) context.Context {
	metaDataMap := map[string]string{}

	requestContext := groupObject.GetRequestContext(ctx)

	// string value
	metaDataMap[string(valueObject.TraceIDHeaderName)] = requestContext.TraceID.GetValue()
	metaDataMap[string(valueObject.ClientIPHeaderName)] = requestContext.ClientIP.GetValue()
	metaDataMap[string(valueObject.UserAgentHeaderName)] = requestContext.UserAgent.GetValue()
	metaDataMap[string(valueObject.UserIDHeaderName)] = requestContext.UserID.GetValue()
	metaDataMap[string(valueObject.AccessTokenHeaderName)] = requestContext.AccessToken.GetValue()
	metaDataMap[string(valueObject.TenantIDHeaderName)] = requestContext.TenantID.GetValue()
	metaDataMap[string(valueObject.LocaleHeaderName)] = requestContext.Locale.GetValue()
	metaDataMap[string(valueObject.TimeZoneHeaderName)] = requestContext.TimeZone.GetValue()

	// int64 value
	metaDataMap[string(valueObject.RequestStartTimeHeaderName)] = requestContext.RequestStartTime.GetString()

	// permissionListを文字列のスライスとして格納
	metaDataMap[string(valueObject.PermissionListHeaderName)] = strings.Join(
		requestContext.PermissionList.GetSliceValue(),
		",",
	)

	md := metadata.New(
		metaDataMap,
	)

	ctx = metadata.NewOutgoingContext(ctx, md)

	return ctx
}

func MetadataToContext(
	ctx context.Context,
) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	newRequestContextArgs := &groupObject.NewRequestContextArgs{}

	// ________________________________
	if len(md.Get(string(valueObject.RequestStartTimeHeaderName))) != 0 {
		requestStartTime, _ := strconv.ParseInt(
			md.Get(string(valueObject.RequestStartTimeHeaderName))[0], 10, 64,
		)
		newRequestContextArgs.RequestStartTime = &requestStartTime
	}

	// ________________________________
	if len(md.Get(string(valueObject.PermissionListHeaderName))) != 0 {
		permissionList := []string{}
		permissionList = append(
			permissionList,
			md.Get(string(valueObject.PermissionListHeaderName))...,
		)
		newRequestContextArgs.PermissionList = permissionList

	}

	// ________________________________
	if len(md.Get(string(valueObject.TraceIDHeaderName))) != 0 {
		value := md.Get(string(valueObject.TraceIDHeaderName))[0]
		newRequestContextArgs.TraceID = &value
	}

	// ________________________________
	if len(md.Get(string(valueObject.ClientIPHeaderName))) != 0 {
		value := md.Get(string(valueObject.ClientIPHeaderName))[0]
		newRequestContextArgs.ClientIP = &value
	}

	// ________________________________
	if len(md.Get(string(valueObject.UserAgentHeaderName))) != 0 {
		value := md.Get(string(valueObject.UserAgentHeaderName))[0]
		newRequestContextArgs.UserAgent = &value
	}

	// ________________________________
	if len(md.Get(string(valueObject.UserIDHeaderName))) != 0 {
		value := md.Get(string(valueObject.UserIDHeaderName))[0]
		newRequestContextArgs.UserID = &value
	}

	// ________________________________
	if len(md.Get(string(valueObject.AccessTokenHeaderName))) != 0 {
		value := md.Get(string(valueObject.AccessTokenHeaderName))[0]
		newRequestContextArgs.AccessToken = &value
	}

	// ________________________________
	if len(md.Get(string(valueObject.TenantIDHeaderName))) != 0 {
		value := md.Get(string(valueObject.TenantIDHeaderName))[0]
		newRequestContextArgs.TenantID = &value
	}

	// ________________________________
	if len(md.Get(string(valueObject.LocaleHeaderName))) != 0 {
		value := md.Get(string(valueObject.LocaleHeaderName))[0]
		newRequestContextArgs.Locale = &value
	}

	// ________________________________
	if len(md.Get(string(valueObject.TimeZoneHeaderName))) != 0 {
		value := md.Get(string(valueObject.TimeZoneHeaderName))[0]
		newRequestContextArgs.TimeZone = &value
	}

	requestContext := groupObject.NewRequestContext(
		ctx,
		newRequestContextArgs,
	)
	if requestContext.GetError() != nil {
		logger.Logging(ctx, requestContext.GetError())
		return ctx
	}

	ctx = context.WithValue(
		ctx,
		groupObject.RequestContextContextName,
		*requestContext,
	)

	// ________________________________
	// logで追跡するために、contextにTraceIDを設定する
	ctx = context.WithValue(
		ctx,
		valueObject.TraceIDContextName,
		requestContext.TraceID.GetValue(),
	)

	return ctx
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (
		interface{},
		error,
	) {
		ctx = MetadataToContext(ctx)

		return handler(ctx, req)
	}
}

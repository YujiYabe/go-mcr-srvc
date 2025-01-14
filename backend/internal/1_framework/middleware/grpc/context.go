package grpc_middleware

import (
	"context"
	"log"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	groupObject "backend/internal/4_domain/group_object"
	primitiveObject "backend/internal/4_domain/primitive_object"
	valueObject "backend/internal/4_domain/value_object"
)

// ------------
func ContextToMetadata(
	ctx context.Context,
) context.Context {
	metaDataMap := map[string]string{}

	requestContext := groupObject.GetRequestContext(ctx)

	// string value
	metaDataMap[string(valueObject.TraceIDMetaName)] = requestContext.TraceID.GetValue()
	metaDataMap[string(valueObject.ClientIPMetaName)] = requestContext.ClientIP.GetValue()
	metaDataMap[string(valueObject.UserAgentMetaName)] = requestContext.UserAgent.GetValue()
	metaDataMap[string(valueObject.UserIDMetaName)] = requestContext.UserID.GetValue()
	metaDataMap[string(valueObject.AccessTokenMetaName)] = requestContext.AccessToken.GetValue()
	metaDataMap[string(valueObject.TenantIDMetaName)] = requestContext.TenantID.GetValue()
	metaDataMap[string(valueObject.LocaleMetaName)] = requestContext.Locale.GetValue()
	metaDataMap[string(valueObject.TimeZoneMetaName)] = requestContext.TimeZone.GetValue()

	// int64 value
	metaDataMap[string(valueObject.RequestStartTimeMetaName)] = requestContext.RequestStartTime.GetString()

	// permissionListを文字列のスライスとして格納
	metaDataMap[string(valueObject.PermissionListMetaName)] = strings.Join(
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

	requestStartTime := int64(0)
	if len(md.Get(string(valueObject.RequestStartTimeMetaName))) != 0 {
		requestStartTime, _ = strconv.ParseInt(
			md.Get(string(valueObject.RequestStartTimeMetaName))[0], 10, 64,
		)
	}
	newRequestContextArgs.RequestStartTime = &requestStartTime

	// ________________________________
	permissionList := []string{}
	permissionList = append(
		permissionList,
		md.Get(string(valueObject.PermissionListMetaName))...,
	)
	newRequestContextArgs.PermissionList = permissionList

	// ________________________________
	newRequestContextArgs.TraceID = primitiveObject.ExtractFirstIndexFromSliceString(
		md.Get(string(valueObject.TraceIDMetaName)),
	)

	// ________________________________
	newRequestContextArgs.ClientIP = primitiveObject.ExtractFirstIndexFromSliceString(
		md.Get(string(valueObject.ClientIPMetaName)),
	)

	// ________________________________
	newRequestContextArgs.UserAgent = primitiveObject.ExtractFirstIndexFromSliceString(
		md.Get(string(valueObject.UserAgentMetaName)),
	)

	// ________________________________
	newRequestContextArgs.UserID = primitiveObject.ExtractFirstIndexFromSliceString(
		md.Get(string(valueObject.UserIDMetaName)),
	)

	// ________________________________
	newRequestContextArgs.AccessToken = primitiveObject.ExtractFirstIndexFromSliceString(
		md.Get(string(valueObject.AccessTokenMetaName)),
	)

	// ________________________________
	newRequestContextArgs.TenantID = primitiveObject.ExtractFirstIndexFromSliceString(
		md.Get(string(valueObject.TenantIDMetaName)),
	)

	// ________________________________
	newRequestContextArgs.Locale = primitiveObject.ExtractFirstIndexFromSliceString(
		md.Get(string(valueObject.LocaleMetaName)),
	)

	// ________________________________
	newRequestContextArgs.TimeZone = primitiveObject.ExtractFirstIndexFromSliceString(
		md.Get(string(valueObject.TimeZoneMetaName)),
	)

	// ________________________________
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

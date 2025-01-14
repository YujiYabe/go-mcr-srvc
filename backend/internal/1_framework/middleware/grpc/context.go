package grpc_middleware

import (
	"context"
	"log"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	groupObject "backend/internal/4_domain/group_object"
	valueObject "backend/internal/4_domain/value_object"
	"backend/pkg"
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

	requestStartTime, _ := strconv.ParseInt(
		md.Get(string(valueObject.RequestStartTimeMetaName))[0], 10, 64,
	)

	permissionList := []string{}
	permissionList = append(
		permissionList,
		md.Get(string(valueObject.PermissionListMetaName))...,
	)

	newRequestContextArgs := &groupObject.NewRequestContextArgs{
		RequestStartTime: &requestStartTime,
		TraceID:          &md.Get(string(valueObject.TraceIDMetaName))[0],
		ClientIP:         &md.Get(string(valueObject.ClientIPMetaName))[0],
		UserAgent:        &md.Get(string(valueObject.UserAgentMetaName))[0],
		UserID:           &md.Get(string(valueObject.UserIDMetaName))[0],
		AccessToken:      &md.Get(string(valueObject.AccessTokenMetaName))[0],
		TenantID:         &md.Get(string(valueObject.TenantIDMetaName))[0],
		Locale:           &md.Get(string(valueObject.LocaleMetaName))[0],
		TimeZone:         &md.Get(string(valueObject.TimeZoneMetaName))[0],
		PermissionList:   permissionList,
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
		pkg.Logging(ctx, "Metadata converted to context")

		return handler(ctx, req)
	}
}

package pubsub_middleware

import (
	"context"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	groupObject "backend/internal/4_domain/group_object"
	primitiveObject "backend/internal/4_domain/primitive_object"
	valueObject "backend/internal/4_domain/value_object"
	"backend/internal/logger"
)

func HeaderToContext(
	headers []kafka.Header,
) (
	ctx context.Context,
) {
	ctx = context.Background()
	newRequestContextArgs := &groupObject.NewRequestContextArgs{}

	// ________________________________
	// pubsubのheaderから、traceIDを取得する
	for _, header := range headers {
		keyString := primitiveObject.ContextKey(header.Key)
		valueString := string(header.Value)

		switch keyString {
		case valueObject.AccessTokenMetaName:
			newRequestContextArgs.AccessToken = &valueString
		case valueObject.ClientIPMetaName:
			newRequestContextArgs.ClientIP = &valueString
		case valueObject.LocaleMetaName:
			newRequestContextArgs.Locale = &valueString

		case valueObject.PermissionListMetaName:
			permissionList := []string{}
			// permissionList = append(
			// 	permissionList,
			// 	md.Get(string(valueObject.PermissionListMetaName))...,
			// )

			newRequestContextArgs.PermissionList = permissionList

		case valueObject.RequestStartTimeMetaName:
			requestStartTime, _ := strconv.ParseInt(valueString, 10, 64)
			newRequestContextArgs.RequestStartTime = &requestStartTime

		case valueObject.TenantIDMetaName:
			newRequestContextArgs.TenantID = &valueString
		case valueObject.TimeZoneMetaName:
			newRequestContextArgs.TimeZone = &valueString
		case valueObject.TraceIDMetaName:
			newRequestContextArgs.TraceID = &valueString
		case valueObject.UserAgentMetaName:
			newRequestContextArgs.UserAgent = &valueString
		case valueObject.UserIDMetaName:
			newRequestContextArgs.UserID = &valueString
		}
	}

	requestContext := groupObject.NewRequestContext(
		ctx,
		newRequestContextArgs,
	)
	if requestContext.GetError() != nil {
		logger.Logging(ctx, requestContext.GetError())
		return
	}

	ctx = context.WithValue(
		ctx,
		groupObject.RequestContextContextName,
		*requestContext,
	)

	return
}

// contextからpubsubのheaderを生成する。MetadataToContextと逆の関数
func ContextToHeader(
	ctx context.Context,
) (
	headers []kafka.Header,
) {
	requestContext := groupObject.GetRequestContext(ctx)
	if requestContext == nil {
		return headers
	}
	headers = []kafka.Header{}

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.AccessTokenMetaName),
			Value: []byte(requestContext.AccessToken.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.ClientIPMetaName),
			Value: []byte(requestContext.AccessToken.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.ClientIPMetaName),
			Value: []byte(requestContext.ClientIP.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.LocaleMetaName),
			Value: []byte(requestContext.Locale.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.RequestStartTimeMetaName),
			Value: []byte(requestContext.RequestStartTime.GetString()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.TenantIDMetaName),
			Value: []byte(requestContext.TenantID.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.TimeZoneMetaName),
			Value: []byte(requestContext.TimeZone.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.TraceIDMetaName),
			Value: []byte(requestContext.TraceID.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.UserAgentMetaName),
			Value: []byte(requestContext.UserAgent.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.UserIDMetaName),
			Value: []byte(requestContext.UserID.GetValue()),
		},
	)

	return

}

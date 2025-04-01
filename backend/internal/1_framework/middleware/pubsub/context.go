package pubsub_middleware

import (
	"context"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	groupObject "backend/internal/4_domain/group_object"
	primitiveObject "backend/internal/4_domain/primitive_object"
	typeObject "backend/internal/4_domain/type_object"
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
		headerKey := primitiveObject.ContextKey(header.Key)
		headerValue := string(header.Value)

		switch headerKey {
		case typeObject.AccessTokenHeaderName:
			newRequestContextArgs.AccessToken = &headerValue
		case typeObject.ClientIPHeaderName:
			newRequestContextArgs.ClientIP = &headerValue
		case typeObject.LocaleHeaderName:
			newRequestContextArgs.Locale = &headerValue
		case typeObject.TenantIDHeaderName:
			newRequestContextArgs.TenantID = &headerValue
		case typeObject.TimeZoneHeaderName:
			newRequestContextArgs.TimeZone = &headerValue
		case typeObject.TraceIDHeaderName:
			newRequestContextArgs.TraceID = &headerValue
		case typeObject.UserAgentHeaderName:
			newRequestContextArgs.UserAgent = &headerValue
		case typeObject.UserIDHeaderName:
			newRequestContextArgs.UserID = &headerValue

		case typeObject.PermissionListHeaderName:
			permissionList := []string{}
			// permissionList = append(
			// 	permissionList,
			// 	md.Get(string(typeObject.PermissionListHeaderName))...,
			// )

			newRequestContextArgs.PermissionList = permissionList

		case typeObject.RequestStartTimeHeaderName:
			requestStartTime, _ := strconv.ParseInt(headerValue, 10, 64)
			newRequestContextArgs.RequestStartTime = &requestStartTime
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

// contextからpubsubのheaderを生成する。 HeaderToContext と逆の関数
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
			Key:   string(typeObject.AccessTokenHeaderName),
			Value: []byte(requestContext.AccessToken.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(typeObject.ClientIPHeaderName),
			Value: []byte(requestContext.AccessToken.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(typeObject.ClientIPHeaderName),
			Value: []byte(requestContext.ClientIP.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(typeObject.LocaleHeaderName),
			Value: []byte(requestContext.Locale.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(typeObject.RequestStartTimeHeaderName),
			Value: []byte(requestContext.RequestStartTime.GetString()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(typeObject.TenantIDHeaderName),
			Value: []byte(requestContext.TenantID.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(typeObject.TimeZoneHeaderName),
			Value: []byte(requestContext.TimeZone.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(typeObject.TraceIDHeaderName),
			Value: []byte(requestContext.TraceID.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(typeObject.UserAgentHeaderName),
			Value: []byte(requestContext.UserAgent.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(typeObject.UserIDHeaderName),
			Value: []byte(requestContext.UserID.GetValue()),
		},
	)

	return

}

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
		headerKey := primitiveObject.ContextKey(header.Key)
		headerValue := string(header.Value)

		switch headerKey {
		case valueObject.AccessTokenHeaderName:
			newRequestContextArgs.AccessToken = &headerValue
		case valueObject.ClientIPHeaderName:
			newRequestContextArgs.ClientIP = &headerValue
		case valueObject.LocaleHeaderName:
			newRequestContextArgs.Locale = &headerValue
		case valueObject.TenantIDHeaderName:
			newRequestContextArgs.TenantID = &headerValue
		case valueObject.TimeZoneHeaderName:
			newRequestContextArgs.TimeZone = &headerValue
		case valueObject.TraceIDHeaderName:
			newRequestContextArgs.TraceID = &headerValue
		case valueObject.UserAgentHeaderName:
			newRequestContextArgs.UserAgent = &headerValue
		case valueObject.UserIDHeaderName:
			newRequestContextArgs.UserID = &headerValue

		case valueObject.PermissionListHeaderName:
			permissionList := []string{}
			// permissionList = append(
			// 	permissionList,
			// 	md.Get(string(valueObject.PermissionListHeaderName))...,
			// )

			newRequestContextArgs.PermissionList = permissionList

		case valueObject.RequestStartTimeHeaderName:
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
			Key:   string(valueObject.AccessTokenHeaderName),
			Value: []byte(requestContext.AccessToken.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.ClientIPHeaderName),
			Value: []byte(requestContext.AccessToken.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.ClientIPHeaderName),
			Value: []byte(requestContext.ClientIP.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.LocaleHeaderName),
			Value: []byte(requestContext.Locale.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.RequestStartTimeHeaderName),
			Value: []byte(requestContext.RequestStartTime.GetString()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.TenantIDHeaderName),
			Value: []byte(requestContext.TenantID.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.TimeZoneHeaderName),
			Value: []byte(requestContext.TimeZone.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.TraceIDHeaderName),
			Value: []byte(requestContext.TraceID.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.UserAgentHeaderName),
			Value: []byte(requestContext.UserAgent.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(valueObject.UserIDHeaderName),
			Value: []byte(requestContext.UserID.GetValue()),
		},
	)

	return

}

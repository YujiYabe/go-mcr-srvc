package pubsub_middleware

import (
	"context"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	domainObject "backend/internal/4_domain/domain_object"
	groupObject "backend/internal/4_domain/group_object"
	primitiveObject "backend/internal/4_domain/primitive_object"
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
		case domainObject.AccessTokenHeaderName:
			newRequestContextArgs.AccessToken = &headerValue
		case domainObject.ClientIPHeaderName:
			newRequestContextArgs.ClientIP = &headerValue
		case domainObject.LocaleHeaderName:
			newRequestContextArgs.Locale = &headerValue
		case domainObject.TenantIDHeaderName:
			newRequestContextArgs.TenantID = &headerValue
		case domainObject.TimeZoneHeaderName:
			newRequestContextArgs.TimeZone = &headerValue
		case domainObject.TraceIDHeaderName:
			newRequestContextArgs.TraceID = &headerValue
		case domainObject.UserAgentHeaderName:
			newRequestContextArgs.UserAgent = &headerValue
		case domainObject.UserIDHeaderName:
			newRequestContextArgs.UserID = &headerValue

		case domainObject.PermissionListHeaderName:
			permissionList := []string{}
			// permissionList = append(
			// 	permissionList,
			// 	md.Get(string(domainObject.PermissionListHeaderName))...,
			// )

			newRequestContextArgs.PermissionList = permissionList

		case domainObject.RequestStartTimeHeaderName:
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
			Key:   string(domainObject.AccessTokenHeaderName),
			Value: []byte(requestContext.AccessToken.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(domainObject.ClientIPHeaderName),
			Value: []byte(requestContext.AccessToken.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(domainObject.ClientIPHeaderName),
			Value: []byte(requestContext.ClientIP.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(domainObject.LocaleHeaderName),
			Value: []byte(requestContext.Locale.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(domainObject.RequestStartTimeHeaderName),
			Value: []byte(requestContext.RequestStartTime.GetString()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(domainObject.TenantIDHeaderName),
			Value: []byte(requestContext.TenantID.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(domainObject.TimeZoneHeaderName),
			Value: []byte(requestContext.TimeZone.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(domainObject.TraceIDHeaderName),
			Value: []byte(requestContext.TraceID.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(domainObject.UserAgentHeaderName),
			Value: []byte(requestContext.UserAgent.GetValue()),
		},
	)

	headers = append(
		headers,
		kafka.Header{
			Key:   string(domainObject.UserIDHeaderName),
			Value: []byte(requestContext.UserID.GetValue()),
		},
	)

	return

}

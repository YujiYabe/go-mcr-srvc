package pubsub_middleware

import (
	"context"
	"log"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	groupObject "backend/internal/4_domain/group_object"
	primitiveObject "backend/internal/4_domain/primitive_object"
	valueObject "backend/internal/4_domain/value_object"
)

func MetadataToContext(msg *kafka.Message) context.Context {
	ctx := context.Background()
	newRequestContextArgs := &groupObject.NewRequestContextArgs{}

	// ________________________________
	// pubsubのheaderから、traceIDを取得する
	for _, header := range msg.Headers {
		keyString := primitiveObject.ContextKey(header.Key)
		valueString := string(header.Value)

		switch keyString {
		case valueObject.AccessTokenMetaName:
			newRequestContextArgs.TraceID = &valueString
		case valueObject.ClientIPMetaName:
			newRequestContextArgs.ClientIP = &valueString
		case valueObject.LocaleContextName:
			newRequestContextArgs.Locale = &valueString

		case valueObject.PermissionListContextName:
			permissionList := []string{}
			// permissionList = append(
			// 	permissionList,
			// 	md.Get(string(valueObject.PermissionListMetaName))...,
			// )

			newRequestContextArgs.PermissionList = permissionList

		case valueObject.RequestStartTimeContextName:
			requestStartTime, _ := strconv.ParseInt(valueString, 10, 64)
			newRequestContextArgs.RequestStartTime = &requestStartTime

		case valueObject.TenantIDContextName:
			newRequestContextArgs.TenantID = &valueString
		case valueObject.TimeZoneContextName:
			newRequestContextArgs.TimeZone = &valueString
		case valueObject.TraceIDContextName:
			newRequestContextArgs.TraceID = &valueString
		case valueObject.UserAgentContextName:
			newRequestContextArgs.UserAgent = &valueString
		case valueObject.UserIDContextName:
			newRequestContextArgs.UserID = &valueString
		}
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

package request_context

import (
	"time"

	primitiveObject "backend/internal/4_domain/primitive_object"
	typeObject "backend/internal/4_domain/type_object"
)

const (
	RequestContextHeaderName  primitiveObject.ContextKey = "request-context"
	RequestContextContextName primitiveObject.ContextKey = "requestContext"
)

var HeaderNameToContextNameMap = map[primitiveObject.ContextKey]primitiveObject.ContextKey{
	typeObject.AccessTokenHeaderName:       typeObject.AccessTokenContextName,
	typeObject.ClientIPHeaderName:          typeObject.ClientIPContextName,
	typeObject.LocaleHeaderName:            typeObject.LocaleContextName,
	typeObject.PermissionListHeaderName:    typeObject.PermissionListContextName,
	typeObject.RequestStartTimeHeaderName:  typeObject.RequestStartTimeContextName,
	typeObject.TenantIDHeaderName:          typeObject.TenantIDContextName,
	typeObject.TimeOutMillSecondHeaderName: typeObject.TimeOutMillSecondContextName,
	typeObject.TimeZoneHeaderName:          typeObject.TimeZoneContextName,
	typeObject.TraceIDHeaderName:           typeObject.TraceIDContextName,
	typeObject.UserAgentHeaderName:         typeObject.UserAgentContextName,
	typeObject.UserIDHeaderName:            typeObject.UserIDContextName,
}

var ContextNameToHeaderNameMap = map[primitiveObject.ContextKey]primitiveObject.ContextKey{
	typeObject.AccessTokenContextName:       typeObject.AccessTokenHeaderName,
	typeObject.ClientIPContextName:          typeObject.ClientIPHeaderName,
	typeObject.LocaleContextName:            typeObject.LocaleHeaderName,
	typeObject.PermissionListContextName:    typeObject.PermissionListHeaderName,
	typeObject.RequestStartTimeContextName:  typeObject.RequestStartTimeHeaderName,
	typeObject.TenantIDContextName:          typeObject.TenantIDHeaderName,
	typeObject.TimeOutMillSecondContextName: typeObject.TimeOutMillSecondHeaderName,
	typeObject.TimeZoneContextName:          typeObject.TimeZoneHeaderName,
	typeObject.TraceIDContextName:           typeObject.TraceIDHeaderName,
	typeObject.UserAgentContextName:         typeObject.UserAgentHeaderName,
	typeObject.UserIDContextName:            typeObject.UserIDHeaderName,
}

type RequestContext struct {
	timeOutMillSecond typeObject.TimeOutMillSecond
	requestStartTime  typeObject.RequestStartTime
	traceID           typeObject.TraceID
	clientIP          typeObject.ClientIP
	userAgent         typeObject.UserAgent
	userID            typeObject.UserID
	accessToken       typeObject.AccessToken
	tenantID          typeObject.TenantID
	locale            typeObject.Locale
	timeZone          typeObject.TimeZone
	permissionList    typeObject.PermissionList
}

type NewRequestContextArgs struct {
	RequestStartTime *int64
	TraceID          *string
	ClientIP         *string
	UserAgent        *string
	UserID           *string
	AccessToken      *string
	TenantID         *string
	Locale           *string
	TimeZone         *string
	PermissionList   []string
}

func NewRequestContext(
	args *NewRequestContextArgs,
) (
	requestContext *RequestContext,
	err error,
) {
	requestContext = &RequestContext{}

	requestContext.requestStartTime, err = typeObject.NewRequestStartTime(args.RequestStartTime)
	if err != nil {
		return nil, err
	}

	requestContext.traceID, err = typeObject.NewTraceID(args.TraceID)
	if err != nil {
		return nil, err
	}

	requestContext.clientIP, err = typeObject.NewClientIP(args.ClientIP)
	if err != nil {
		return nil, err
	}

	requestContext.userAgent, err = typeObject.NewUserAgent(args.UserAgent)
	if err != nil {
		return nil, err
	}

	requestContext.locale, err = typeObject.NewLocale(args.Locale)
	if err != nil {
		return nil, err
	}

	requestContext.timeZone, err = typeObject.NewTimeZone(args.TimeZone)
	if err != nil {
		return nil, err
	}

	requestContext.userID, err = typeObject.NewUserID(args.UserID)
	if err != nil {
		return nil, err
	}

	requestContext.accessToken, err = typeObject.NewAccessToken(args.AccessToken)
	if err != nil {
		return nil, err
	}

	requestContext.tenantID, err = typeObject.NewTenantID(args.TenantID)
	if err != nil {
		return nil, err
	}

	requestContext.permissionList, err = typeObject.NewPermissionList(args.PermissionList)
	if err != nil {
		return nil, err
	}

	requestStartTime := requestContext.requestStartTime
	requestEndTime := time.UnixMilli(requestStartTime.GetValue()).Add(typeObject.TimeOutMillSecondValue * time.Second).UnixMilli()
	timeoutMillSecond := requestEndTime - time.Now().UnixMilli()

	requestContext.timeOutMillSecond, err = typeObject.NewTimeOutMillSecond(&timeoutMillSecond)
	if err != nil {
		return nil, err
	}

	return
}

func (receiver RequestContext) TimeOutMillSecond() (
	timeOutMillSecond typeObject.TimeOutMillSecond,
) {
	return receiver.timeOutMillSecond
}

func (receiver RequestContext) RequestStartTime() (
	requestStartTime typeObject.RequestStartTime,
) {
	return receiver.requestStartTime
}

func (receiver RequestContext) TraceID() (
	traceID typeObject.TraceID,
) {
	return receiver.traceID
}

func (receiver RequestContext) ClientIP() (
	clientIP typeObject.ClientIP,
) {
	return receiver.clientIP
}

func (receiver RequestContext) UserAgent() (
	userAgent typeObject.UserAgent,
) {
	return receiver.userAgent
}

func (receiver RequestContext) UserID() (
	userID typeObject.UserID,
) {
	return receiver.userID
}

func (receiver RequestContext) AccessToken() (
	accessToken typeObject.AccessToken,
) {
	return receiver.accessToken
}

func (receiver RequestContext) TenantID() (
	tenantID typeObject.TenantID,
) {
	return receiver.tenantID
}

func (receiver RequestContext) Locale() (
	locale typeObject.Locale,
) {
	return receiver.locale
}

func (receiver RequestContext) TimeZone() (
	timeZone typeObject.TimeZone,
) {
	return receiver.timeZone
}

func (receiver RequestContext) PermissionList() (
	permissionList typeObject.PermissionList,
) {
	return receiver.permissionList
}

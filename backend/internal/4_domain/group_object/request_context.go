package group_object

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
	timeOutMillSecond typeObject.TimeOutMillSecond // RequestStartTimeからの経過時間を格納
	requestStartTime  typeObject.RequestStartTime  // httpかgrpcのリクエスト開始時間を格納
	traceID           typeObject.TraceID           // uuidを格納
	clientIP          typeObject.ClientIP          // httpアクセス元のIPを格納
	userAgent         typeObject.UserAgent         // httpアクセス元のUserAgentを格納
	userID            typeObject.UserID            // 認証ユーザーIDを格納
	accessToken       typeObject.AccessToken       // 認証トークンを格納
	tenantID          typeObject.TenantID          // 所属テナントIDを格納
	locale            typeObject.Locale            // ロケールを格納
	timeZone          typeObject.TimeZone          // タイムゾーンを格納
	permissionList    typeObject.PermissionList    // ユーザー権限を格納
}

type NewRequestContextArgs struct {
	RequestStartTime *int64   //
	TraceID          *string  //
	ClientIP         *string  //
	UserAgent        *string  //
	UserID           *string  //
	AccessToken      *string  //
	TenantID         *string  //
	Locale           *string  //
	TimeZone         *string  //
	PermissionList   []string //
}

func NewRequestContext(
	args *NewRequestContextArgs,
) (
	requestContext *RequestContext,
	err error,
) {
	requestContext = &RequestContext{}

	// ______________________________________
	requestContext.requestStartTime, err = typeObject.NewRequestStartTime(args.RequestStartTime)
	if err != nil {
		return nil, err
	}

	// ______________________________________
	requestContext.traceID, err = typeObject.NewTraceID(args.TraceID)
	if err != nil {
		return nil, err
	}

	// ______________________________________
	requestContext.clientIP, err = typeObject.NewClientIP(args.ClientIP)
	if err != nil {
		return nil, err
	}

	// ______________________________________
	requestContext.userAgent, err = typeObject.NewUserAgent(args.UserAgent)
	if err != nil {
		return nil, err
	}

	// ______________________________________
	requestContext.locale, err = typeObject.NewLocale(args.Locale)
	if err != nil {
		return nil, err
	}

	// ______________________________________
	requestContext.timeZone, err = typeObject.NewTimeZone(args.TimeZone)
	if err != nil {
		return nil, err
	}

	// ______________________________________
	requestContext.userID, err = typeObject.NewUserID(args.UserID)
	if err != nil {
		return nil, err
	}

	// ______________________________________
	requestContext.accessToken, err = typeObject.NewAccessToken(args.AccessToken)
	if err != nil {
		return nil, err
	}

	// ______________________________________
	requestContext.tenantID, err = typeObject.NewTenantID(args.TenantID)
	if err != nil {
		return nil, err
	}

	requestContext.permissionList, err = typeObject.NewPermissionList(args.PermissionList)
	if err != nil {
		return nil, err
	}

	// ______________________________________
	requestStartTime := requestContext.requestStartTime
	requestEndTime := time.UnixMilli(requestStartTime.GetValue()).Add(typeObject.TimeOutMillSecondValue * time.Second).UnixMilli()
	timeoutMillSecond := requestEndTime - time.Now().UnixMilli()

	requestContext.timeOutMillSecond, err = typeObject.NewTimeOutMillSecond(&timeoutMillSecond)
	if err != nil {
		return nil, err
	}

	return
}

func (receiver *RequestContext) TimeOutMillSecond() *typeObject.TimeOutMillSecond {
	return &receiver.timeOutMillSecond
}

func (receiver *RequestContext) RequestStartTime() *typeObject.RequestStartTime {
	return &receiver.requestStartTime
}

func (receiver *RequestContext) TraceID() *typeObject.TraceID {
	return &receiver.traceID
}

func (receiver *RequestContext) ClientIP() *typeObject.ClientIP {
	return &receiver.clientIP
}

func (receiver *RequestContext) UserAgent() *typeObject.UserAgent {
	return &receiver.userAgent
}

func (receiver *RequestContext) UserID() *typeObject.UserID {
	return &receiver.userID
}

func (receiver *RequestContext) AccessToken() *typeObject.AccessToken {
	return &receiver.accessToken
}

func (receiver *RequestContext) TenantID() *typeObject.TenantID {
	return &receiver.tenantID
}

func (receiver *RequestContext) Locale() *typeObject.Locale {
	return &receiver.locale
}

func (receiver *RequestContext) TimeZone() *typeObject.TimeZone {
	return &receiver.timeZone
}

func (receiver *RequestContext) PermissionList() *typeObject.PermissionList {
	return &receiver.permissionList
}

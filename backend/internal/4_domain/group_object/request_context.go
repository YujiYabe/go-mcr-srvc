package group_object

import (
	"context"
	"time"

	primitiveObject "backend/internal/4_domain/primitive_object"
	domainObject "backend/internal/4_domain/type_object"
	"backend/internal/logger"
)

const (
	RequestContextHeaderName  primitiveObject.ContextKey = "request-context"
	RequestContextContextName primitiveObject.ContextKey = "requestContext"
)

var HeaderNameToContextNameMap = map[primitiveObject.ContextKey]primitiveObject.ContextKey{
	domainObject.AccessTokenHeaderName:       domainObject.AccessTokenContextName,
	domainObject.ClientIPHeaderName:          domainObject.ClientIPContextName,
	domainObject.LocaleHeaderName:            domainObject.LocaleContextName,
	domainObject.PermissionListHeaderName:    domainObject.PermissionListContextName,
	domainObject.RequestStartTimeHeaderName:  domainObject.RequestStartTimeContextName,
	domainObject.TenantIDHeaderName:          domainObject.TenantIDContextName,
	domainObject.TimeOutMillSecondHeaderName: domainObject.TimeOutMillSecondContextName,
	domainObject.TimeZoneHeaderName:          domainObject.TimeZoneContextName,
	domainObject.TraceIDHeaderName:           domainObject.TraceIDContextName,
	domainObject.UserAgentHeaderName:         domainObject.UserAgentContextName,
	domainObject.UserIDHeaderName:            domainObject.UserIDContextName,
}

var ContextNameToHeaderNameMap = map[primitiveObject.ContextKey]primitiveObject.ContextKey{
	domainObject.AccessTokenContextName:       domainObject.AccessTokenHeaderName,
	domainObject.ClientIPContextName:          domainObject.ClientIPHeaderName,
	domainObject.LocaleContextName:            domainObject.LocaleHeaderName,
	domainObject.PermissionListContextName:    domainObject.PermissionListHeaderName,
	domainObject.RequestStartTimeContextName:  domainObject.RequestStartTimeHeaderName,
	domainObject.TenantIDContextName:          domainObject.TenantIDHeaderName,
	domainObject.TimeOutMillSecondContextName: domainObject.TimeOutMillSecondHeaderName,
	domainObject.TimeZoneContextName:          domainObject.TimeZoneHeaderName,
	domainObject.TraceIDContextName:           domainObject.TraceIDHeaderName,
	domainObject.UserAgentContextName:         domainObject.UserAgentHeaderName,
	domainObject.UserIDContextName:            domainObject.UserIDHeaderName,
}

type RequestContext struct {
	err               error                          // contextに含める構造体作成時に発生したエラーを格納
	TimeOutMillSecond domainObject.TimeOutMillSecond // RequestStartTimeからの経過時間を格納
	RequestStartTime  domainObject.RequestStartTime  // httpかgrpcのリクエスト開始時間を格納
	TraceID           domainObject.TraceID           // uuidを格納
	ClientIP          domainObject.ClientIP          // httpアクセス元のIPを格納
	UserAgent         domainObject.UserAgent         // httpアクセス元のUserAgentを格納
	UserID            domainObject.UserID            // 認証ユーザーIDを格納
	AccessToken       domainObject.AccessToken       // 認証トークンを格納
	TenantID          domainObject.TenantID          // 所属テナントIDを格納
	Locale            domainObject.Locale            // ロケールを格納
	TimeZone          domainObject.TimeZone          // タイムゾーンを格納
	PermissionList    domainObject.PermissionList    // ユーザー権限を格納
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
	ctx context.Context,
	args *NewRequestContextArgs,
) (
	requestContext *RequestContext,
) {
	requestContext = &RequestContext{}

	// ______________________________________
	requestContext.RequestStartTime = domainObject.NewRequestStartTime(ctx, args.RequestStartTime)
	if requestContext.RequestStartTime.GetError() != nil {
		logger.Logging(ctx, requestContext.RequestStartTime.GetError())
		requestContext.SetError(ctx, requestContext.RequestStartTime.GetError())
		return
	}

	// ______________________________________
	requestContext.TraceID = domainObject.NewTraceID(ctx, args.TraceID)
	if requestContext.TraceID.GetError() != nil {
		logger.Logging(ctx, requestContext.TraceID.GetError())
		requestContext.SetError(ctx, requestContext.TraceID.GetError())
		return
	}

	// ______________________________________
	requestContext.ClientIP = domainObject.NewClientIP(ctx, args.ClientIP)
	if requestContext.ClientIP.GetError() != nil {
		logger.Logging(ctx, requestContext.ClientIP.GetError())
		requestContext.SetError(ctx, requestContext.ClientIP.GetError())
		return
	}

	// ______________________________________
	requestContext.UserAgent = domainObject.NewUserAgent(ctx, args.UserAgent)
	if requestContext.UserAgent.GetError() != nil {
		logger.Logging(ctx, requestContext.UserAgent.GetError())
		requestContext.SetError(ctx, requestContext.UserAgent.GetError())
		return
	}

	// ______________________________________
	requestContext.Locale = domainObject.NewLocale(ctx, args.Locale)
	if requestContext.Locale.GetError() != nil {
		logger.Logging(ctx, requestContext.Locale.GetError())
		requestContext.SetError(ctx, requestContext.Locale.GetError())
		return
	}

	// ______________________________________
	requestContext.TimeZone = domainObject.NewTimeZone(ctx, args.TimeZone)
	if requestContext.TimeZone.GetError() != nil {
		logger.Logging(ctx, requestContext.TimeZone.GetError())
		requestContext.SetError(ctx, requestContext.TimeZone.GetError())
		return
	}

	// ______________________________________
	requestContext.UserID = domainObject.NewUserID(ctx, args.UserID)
	if requestContext.UserID.GetError() != nil {
		logger.Logging(ctx, requestContext.UserID.GetError())
		requestContext.SetError(ctx, requestContext.UserID.GetError())
		return
	}

	// ______________________________________
	requestContext.AccessToken = domainObject.NewAccessToken(ctx, args.AccessToken)
	if requestContext.AccessToken.GetError() != nil {
		logger.Logging(ctx, requestContext.AccessToken.GetError())
		requestContext.SetError(ctx, requestContext.AccessToken.GetError())
		return
	}

	// ______________________________________
	requestContext.TenantID = domainObject.NewTenantID(ctx, args.TenantID)
	if requestContext.TenantID.GetError() != nil {
		logger.Logging(ctx, requestContext.TenantID.GetError())
		requestContext.SetError(ctx, requestContext.TenantID.GetError())
		return
	}

	// ______________________________________
	requestStartTime := requestContext.RequestStartTime
	requestEndTime := time.UnixMilli(requestStartTime.GetValue()).Add(domainObject.TimeOutMillSecondValue * time.Second).UnixMilli()
	timeoutMillSecond := requestEndTime - time.Now().UnixMilli()

	requestContext.TimeOutMillSecond = domainObject.NewTimeOutMillSecond(ctx, &timeoutMillSecond)
	if requestContext.TimeOutMillSecond.GetError() != nil {
		logger.Logging(ctx, requestContext.TimeOutMillSecond.GetError())
		requestContext.SetError(ctx, requestContext.TimeOutMillSecond.GetError())
		return
	}

	return
}

func (receiver *RequestContext) GetError() error {
	return receiver.err
}

func (receiver *RequestContext) SetError(
	ctx context.Context,
	err error,
) {
	if receiver.err == nil {
		receiver.err = err
		logger.Logging(ctx, receiver.GetError())
	}
}

func GetRequestContext(
	ctx context.Context,
) (
	value *RequestContext,
) {
	requestContext, ok := ctx.Value(RequestContextContextName).(RequestContext)
	if ok {
		value = &requestContext
	}

	return
}

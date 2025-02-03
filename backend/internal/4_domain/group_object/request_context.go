package group_object

import (
	"context"
	"time"

	primitiveObject "backend/internal/4_domain/primitive_object"
	valueObject "backend/internal/4_domain/value_object"
	"backend/internal/logger"
)

const (
	RequestContextHeaderName  primitiveObject.ContextKey = "request-context"
	RequestContextContextName primitiveObject.ContextKey = "requestContext"
)

var HeaderNameToContextNameMap = map[primitiveObject.ContextKey]primitiveObject.ContextKey{
	valueObject.AccessTokenHeaderName:       valueObject.AccessTokenContextName,
	valueObject.ClientIPHeaderName:          valueObject.ClientIPContextName,
	valueObject.LocaleHeaderName:            valueObject.LocaleContextName,
	valueObject.PermissionListHeaderName:    valueObject.PermissionListContextName,
	valueObject.RequestStartTimeHeaderName:  valueObject.RequestStartTimeContextName,
	valueObject.TenantIDHeaderName:          valueObject.TenantIDContextName,
	valueObject.TimeOutMillSecondHeaderName: valueObject.TimeOutMillSecondContextName,
	valueObject.TimeZoneHeaderName:          valueObject.TimeZoneContextName,
	valueObject.TraceIDHeaderName:           valueObject.TraceIDContextName,
	valueObject.UserAgentHeaderName:         valueObject.UserAgentContextName,
	valueObject.UserIDHeaderName:            valueObject.UserIDContextName,
}

var ContextNameToHeaderNameMap = map[primitiveObject.ContextKey]primitiveObject.ContextKey{
	valueObject.AccessTokenContextName:       valueObject.AccessTokenHeaderName,
	valueObject.ClientIPContextName:          valueObject.ClientIPHeaderName,
	valueObject.LocaleContextName:            valueObject.LocaleHeaderName,
	valueObject.PermissionListContextName:    valueObject.PermissionListHeaderName,
	valueObject.RequestStartTimeContextName:  valueObject.RequestStartTimeHeaderName,
	valueObject.TenantIDContextName:          valueObject.TenantIDHeaderName,
	valueObject.TimeOutMillSecondContextName: valueObject.TimeOutMillSecondHeaderName,
	valueObject.TimeZoneContextName:          valueObject.TimeZoneHeaderName,
	valueObject.TraceIDContextName:           valueObject.TraceIDHeaderName,
	valueObject.UserAgentContextName:         valueObject.UserAgentHeaderName,
	valueObject.UserIDContextName:            valueObject.UserIDHeaderName,
}

type RequestContext struct {
	err               error                         // contextに含める構造体作成時に発生したエラーを格納
	TimeOutMillSecond valueObject.TimeOutMillSecond // RequestStartTimeからの経過時間を格納
	RequestStartTime  valueObject.RequestStartTime  // httpかgrpcのリクエスト開始時間を格納
	TraceID           valueObject.TraceID           // uuidを格納
	ClientIP          valueObject.ClientIP          // httpアクセス元のIPを格納
	UserAgent         valueObject.UserAgent         // httpアクセス元のUserAgentを格納
	UserID            valueObject.UserID            // 認証ユーザーIDを格納
	AccessToken       valueObject.AccessToken       // 認証トークンを格納
	TenantID          valueObject.TenantID          // 所属テナントIDを格納
	Locale            valueObject.Locale            // ロケールを格納
	TimeZone          valueObject.TimeZone          // タイムゾーンを格納
	PermissionList    valueObject.PermissionList    // ユーザー権限を格納
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
	requestContext.RequestStartTime = valueObject.NewRequestStartTime(ctx, args.RequestStartTime)
	if requestContext.RequestStartTime.GetError() != nil {
		logger.Logging(ctx, requestContext.RequestStartTime.GetError())
		requestContext.SetError(ctx, requestContext.RequestStartTime.GetError())
		return
	}

	// ______________________________________
	requestContext.TraceID = valueObject.NewTraceID(ctx, args.TraceID)
	if requestContext.TraceID.GetError() != nil {
		logger.Logging(ctx, requestContext.TraceID.GetError())
		requestContext.SetError(ctx, requestContext.TraceID.GetError())
		return
	}

	// ______________________________________
	requestContext.ClientIP = valueObject.NewClientIP(ctx, args.ClientIP)
	if requestContext.ClientIP.GetError() != nil {
		logger.Logging(ctx, requestContext.ClientIP.GetError())
		requestContext.SetError(ctx, requestContext.ClientIP.GetError())
		return
	}

	// ______________________________________
	requestContext.UserAgent = valueObject.NewUserAgent(ctx, args.UserAgent)
	if requestContext.UserAgent.GetError() != nil {
		logger.Logging(ctx, requestContext.UserAgent.GetError())
		requestContext.SetError(ctx, requestContext.UserAgent.GetError())
		return
	}

	// ______________________________________
	requestContext.Locale = valueObject.NewLocale(ctx, args.Locale)
	if requestContext.Locale.GetError() != nil {
		logger.Logging(ctx, requestContext.Locale.GetError())
		requestContext.SetError(ctx, requestContext.Locale.GetError())
		return
	}

	// ______________________________________
	requestContext.TimeZone = valueObject.NewTimeZone(ctx, args.TimeZone)
	if requestContext.TimeZone.GetError() != nil {
		logger.Logging(ctx, requestContext.TimeZone.GetError())
		requestContext.SetError(ctx, requestContext.TimeZone.GetError())
		return
	}

	// ______________________________________
	requestContext.UserID = valueObject.NewUserID(ctx, args.UserID)
	if requestContext.UserID.GetError() != nil {
		logger.Logging(ctx, requestContext.UserID.GetError())
		requestContext.SetError(ctx, requestContext.UserID.GetError())
		return
	}

	// ______________________________________
	requestContext.AccessToken = valueObject.NewAccessToken(ctx, args.AccessToken)
	if requestContext.AccessToken.GetError() != nil {
		logger.Logging(ctx, requestContext.AccessToken.GetError())
		requestContext.SetError(ctx, requestContext.AccessToken.GetError())
		return
	}

	// ______________________________________
	requestContext.TenantID = valueObject.NewTenantID(ctx, args.TenantID)
	if requestContext.TenantID.GetError() != nil {
		logger.Logging(ctx, requestContext.TenantID.GetError())
		requestContext.SetError(ctx, requestContext.TenantID.GetError())
		return
	}

	// ______________________________________
	requestStartTime := requestContext.RequestStartTime
	requestEndTime := time.UnixMilli(requestStartTime.GetValue()).Add(valueObject.TimeOutMillSecondValue * time.Second).UnixMilli()
	timeoutMillSecond := requestEndTime - time.Now().UnixMilli()

	requestContext.TimeOutMillSecond = valueObject.NewTimeOutMillSecond(ctx, &timeoutMillSecond)
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

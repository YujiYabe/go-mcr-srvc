package group_object

import (
	"context"
	"time"

	primitiveObject "backend/internal/4_domain/primitive_object"
	typeObject "backend/internal/4_domain/type_object"
	"backend/internal/logger"
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
	err               error                        // contextに含める構造体作成時に発生したエラーを格納
	TimeOutMillSecond typeObject.TimeOutMillSecond // RequestStartTimeからの経過時間を格納
	RequestStartTime  typeObject.RequestStartTime  // httpかgrpcのリクエスト開始時間を格納
	TraceID           typeObject.TraceID           // uuidを格納
	ClientIP          typeObject.ClientIP          // httpアクセス元のIPを格納
	UserAgent         typeObject.UserAgent         // httpアクセス元のUserAgentを格納
	UserID            typeObject.UserID            // 認証ユーザーIDを格納
	AccessToken       typeObject.AccessToken       // 認証トークンを格納
	TenantID          typeObject.TenantID          // 所属テナントIDを格納
	Locale            typeObject.Locale            // ロケールを格納
	TimeZone          typeObject.TimeZone          // タイムゾーンを格納
	PermissionList    typeObject.PermissionList    // ユーザー権限を格納
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
	requestContext.RequestStartTime = typeObject.NewRequestStartTime(ctx, args.RequestStartTime)
	if requestContext.RequestStartTime.GetError() != nil {
		logger.Logging(ctx, requestContext.RequestStartTime.GetError())
		requestContext.SetError(ctx, requestContext.RequestStartTime.GetError())
		return
	}

	// ______________________________________
	requestContext.TraceID = typeObject.NewTraceID(ctx, args.TraceID)
	if requestContext.TraceID.GetError() != nil {
		logger.Logging(ctx, requestContext.TraceID.GetError())
		requestContext.SetError(ctx, requestContext.TraceID.GetError())
		return
	}

	// ______________________________________
	requestContext.ClientIP = typeObject.NewClientIP(ctx, args.ClientIP)
	if requestContext.ClientIP.GetError() != nil {
		logger.Logging(ctx, requestContext.ClientIP.GetError())
		requestContext.SetError(ctx, requestContext.ClientIP.GetError())
		return
	}

	// ______________________________________
	requestContext.UserAgent = typeObject.NewUserAgent(ctx, args.UserAgent)
	if requestContext.UserAgent.GetError() != nil {
		logger.Logging(ctx, requestContext.UserAgent.GetError())
		requestContext.SetError(ctx, requestContext.UserAgent.GetError())
		return
	}

	// ______________________________________
	requestContext.Locale = typeObject.NewLocale(ctx, args.Locale)
	if requestContext.Locale.GetError() != nil {
		logger.Logging(ctx, requestContext.Locale.GetError())
		requestContext.SetError(ctx, requestContext.Locale.GetError())
		return
	}

	// ______________________________________
	requestContext.TimeZone = typeObject.NewTimeZone(ctx, args.TimeZone)
	if requestContext.TimeZone.GetError() != nil {
		logger.Logging(ctx, requestContext.TimeZone.GetError())
		requestContext.SetError(ctx, requestContext.TimeZone.GetError())
		return
	}

	// ______________________________________
	requestContext.UserID = typeObject.NewUserID(ctx, args.UserID)
	if requestContext.UserID.GetError() != nil {
		logger.Logging(ctx, requestContext.UserID.GetError())
		requestContext.SetError(ctx, requestContext.UserID.GetError())
		return
	}

	// ______________________________________
	requestContext.AccessToken = typeObject.NewAccessToken(ctx, args.AccessToken)
	if requestContext.AccessToken.GetError() != nil {
		logger.Logging(ctx, requestContext.AccessToken.GetError())
		requestContext.SetError(ctx, requestContext.AccessToken.GetError())
		return
	}

	// ______________________________________
	requestContext.TenantID = typeObject.NewTenantID(ctx, args.TenantID)
	if requestContext.TenantID.GetError() != nil {
		logger.Logging(ctx, requestContext.TenantID.GetError())
		requestContext.SetError(ctx, requestContext.TenantID.GetError())
		return
	}

	// ______________________________________
	requestStartTime := requestContext.RequestStartTime
	requestEndTime := time.UnixMilli(requestStartTime.GetValue()).Add(typeObject.TimeOutMillSecondValue * time.Second).UnixMilli()
	timeoutMillSecond := requestEndTime - time.Now().UnixMilli()

	requestContext.TimeOutMillSecond = typeObject.NewTimeOutMillSecond(ctx, &timeoutMillSecond)
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

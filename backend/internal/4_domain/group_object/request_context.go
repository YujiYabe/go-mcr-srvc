package group_object

import (
	"context"
	"time"

	primitiveObject "backend/internal/4_domain/primitive_object"
	valueObject "backend/internal/4_domain/value_object"
	"backend/internal/logger"
)

const (
	RequestContextMetaName    primitiveObject.ContextKey = "request-context"
	RequestContextContextName primitiveObject.ContextKey = "requestContext"
)

type RequestContext struct {
	err              error
	RequestStartTime valueObject.RequestStartTime
	TraceID          valueObject.TraceID
	ClientIP         valueObject.ClientIP
	UserAgent        valueObject.UserAgent
	UserID           valueObject.UserID
	AccessToken      valueObject.AccessToken
	TenantID         valueObject.TenantID
	Locale           valueObject.Locale
	TimeZone         valueObject.TimeZone
	TimeOutSecond    valueObject.TimeOutSecond
	PermissionList   valueObject.PermissionList
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
	ctx context.Context,
	args *NewRequestContextArgs,
) (
	requestContext *RequestContext,
) {
	requestContext = &RequestContext{}

	requestContext.RequestStartTime = valueObject.NewRequestStartTime(ctx, args.RequestStartTime)
	if requestContext.RequestStartTime.GetError() != nil {
		logger.Logging(ctx, requestContext.RequestStartTime.GetError())
		requestContext.SetError(ctx, requestContext.RequestStartTime.GetError())
		return
	}

	requestContext.TraceID = valueObject.NewTraceID(ctx, args.TraceID)
	if requestContext.TraceID.GetError() != nil {
		logger.Logging(ctx, requestContext.TraceID.GetError())
		requestContext.SetError(ctx, requestContext.TraceID.GetError())
		return
	}

	requestContext.ClientIP = valueObject.NewClientIP(ctx, args.ClientIP)
	if requestContext.ClientIP.GetError() != nil {
		logger.Logging(ctx, requestContext.ClientIP.GetError())
		requestContext.SetError(ctx, requestContext.ClientIP.GetError())
		return
	}

	requestContext.UserAgent = valueObject.NewUserAgent(ctx, args.UserAgent)
	if requestContext.UserAgent.GetError() != nil {
		logger.Logging(ctx, requestContext.UserAgent.GetError())
		requestContext.SetError(ctx, requestContext.UserAgent.GetError())
		return
	}

	requestContext.Locale = valueObject.NewLocale(ctx, args.Locale)
	if requestContext.Locale.GetError() != nil {
		logger.Logging(ctx, requestContext.Locale.GetError())
		requestContext.SetError(ctx, requestContext.Locale.GetError())
		return
	}

	requestContext.TimeZone = valueObject.NewTimeZone(ctx, args.TimeZone)
	if requestContext.TimeZone.GetError() != nil {
		logger.Logging(ctx, requestContext.TimeZone.GetError())
		requestContext.SetError(ctx, requestContext.TimeZone.GetError())
		return
	}

	requestContext.UserID = valueObject.NewUserID(ctx, args.UserID)
	if requestContext.UserID.GetError() != nil {
		logger.Logging(ctx, requestContext.UserID.GetError())
		requestContext.SetError(ctx, requestContext.UserID.GetError())
		return
	}

	requestContext.AccessToken = valueObject.NewAccessToken(ctx, args.AccessToken)
	if requestContext.AccessToken.GetError() != nil {
		logger.Logging(ctx, requestContext.AccessToken.GetError())
		requestContext.SetError(ctx, requestContext.AccessToken.GetError())
		return
	}

	requestContext.TenantID = valueObject.NewTenantID(ctx, args.TenantID)
	if requestContext.TenantID.GetError() != nil {
		logger.Logging(ctx, requestContext.TenantID.GetError())
		requestContext.SetError(ctx, requestContext.TenantID.GetError())
		return
	}

	requestStartTime := requestContext.RequestStartTime
	currentTimestamp := time.Now().UnixMilli()
	requestEndTime := time.UnixMilli(requestStartTime.GetValue()).Add(valueObject.TimeOutSecondValue * time.Second).UnixMilli()
	timeoutSecond := requestEndTime - currentTimestamp

	requestContext.TimeOutSecond = valueObject.NewTimeOutSecond(ctx, &timeoutSecond)
	if requestContext.TimeOutSecond.GetError() != nil {
		logger.Logging(ctx, requestContext.TimeOutSecond.GetError())
		requestContext.SetError(ctx, requestContext.TimeOutSecond.GetError())
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

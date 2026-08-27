package pubsub_middleware

import (
	"context"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	middlewareRequestContext "backend/internal/1_framework/middleware/request_context"
	typeObject "backend/internal/4_domain/type_object"
)

func TestContextToHeaderAndHeaderToContextRoundTrip(t *testing.T) {
	requestContext := newRequestContextForTest(t)
	ctx := context.WithValue(
		context.Background(),
		middlewareRequestContext.RequestContextContextName,
		*requestContext,
	)

	roundTripCtx := HeaderToContext(context.Background(), ContextToHeader(ctx))

	assertRequestContextForTest(roundTripCtx, t)
}

func TestContextToHeaderIncludesPermissionList(t *testing.T) {
	requestContext := newRequestContextForTest(t)
	ctx := context.WithValue(
		context.Background(),
		middlewareRequestContext.RequestContextContextName,
		*requestContext,
	)

	headers := ContextToHeader(ctx)
	got := findHeaderValue(headers, string(typeObject.PermissionListHeaderName))

	if got != typeObject.PermissionUserRead+","+typeObject.PermissionUserWrite {
		t.Fatalf("unexpected permissions header: %q", got)
	}
}

func TestHeaderToContextRestoresPermissionList(t *testing.T) {
	ctx := HeaderToContext(context.Background(), []kafka.Header{
		{
			Key: string(typeObject.PermissionListHeaderName),
			Value: []byte(
				typeObject.PermissionUserRead + "," + typeObject.PermissionUserWrite,
			),
		},
	})

	requestContext := middlewareRequestContext.GetRequestContext(ctx)
	if requestContext == nil {
		t.Fatal("expected request context")
	}

	permissionList := requestContext.PermissionList()
	if !permissionList.CanReadUser() {
		t.Fatal("expected user:read permission")
	}
	if !permissionList.CanWriteUser() {
		t.Fatal("expected user:write permission")
	}
	if permissionList.Count() != 2 {
		t.Fatalf("expected 2 permissions, got %d", permissionList.Count())
	}
}

func TestContextToHeaderWithoutRequestContextReturnsEmptyHeaders(t *testing.T) {
	headers := ContextToHeader(context.Background())

	if len(headers) != 0 {
		t.Fatalf("expected empty headers, got %d", len(headers))
	}
}

func newRequestContextForTest(t *testing.T) *middlewareRequestContext.RequestContext {
	t.Helper()

	traceID := "123e4567-e89b-12d3-a456-426614174000"
	clientIP := "192.0.2.1"
	userAgent := "go-test"
	userID := "user-1"
	accessToken := "access-token"
	tenantID := "tenant-1"
	locale := "ja-JP"
	timeZone := "AsiaTokyo"

	requestContext, err := middlewareRequestContext.NewRequestContext(
		&middlewareRequestContext.NewRequestContextArgs{
			TraceID:     &traceID,
			ClientIP:    &clientIP,
			UserAgent:   &userAgent,
			UserID:      &userID,
			AccessToken: &accessToken,
			TenantID:    &tenantID,
			Locale:      &locale,
			TimeZone:    &timeZone,
			PermissionList: []string{
				typeObject.PermissionUserRead,
				typeObject.PermissionUserWrite,
			},
		},
	)
	if err != nil {
		t.Fatalf("new request context: %v", err)
	}

	return requestContext
}

func assertRequestContextForTest(ctx context.Context, t *testing.T) {
	t.Helper()

	requestContext := middlewareRequestContext.GetRequestContext(ctx)
	if requestContext == nil {
		t.Fatal("expected request context")
	}

	if got := requestContext.TraceID().GetValue(); got != "123e4567-e89b-12d3-a456-426614174000" {
		t.Fatalf("unexpected trace id: %q", got)
	}
	if got := requestContext.ClientIP().GetValue(); got != "192.0.2.1" {
		t.Fatalf("unexpected client ip: %q", got)
	}
	if got := requestContext.AccessToken().GetValue(); got != "access-token" {
		t.Fatalf("unexpected access token: %q", got)
	}
	if got := requestContext.UserID().GetValue(); got != "user-1" {
		t.Fatalf("unexpected user id: %q", got)
	}

	permissionList := requestContext.PermissionList()
	if !permissionList.CanReadUser() {
		t.Fatal("expected user:read permission")
	}
	if !permissionList.CanWriteUser() {
		t.Fatal("expected user:write permission")
	}
	if permissionList.Count() != 2 {
		t.Fatalf("expected 2 permissions, got %d", permissionList.Count())
	}
}

func findHeaderValue(headers []kafka.Header, key string) string {
	for _, header := range headers {
		if header.Key == key {
			return string(header.Value)
		}
	}

	return ""
}

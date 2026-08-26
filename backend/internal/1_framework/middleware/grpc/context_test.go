package grpc_middleware

import (
	"context"
	"testing"

	"google.golang.org/grpc/metadata"

	requestContextMiddleware "backend/internal/1_framework/middleware/request_context"
	typeObject "backend/internal/4_domain/type_object"
)

func TestContextToMetadataAndMetadataToContextRoundTrip(t *testing.T) {
	outgoingCtx := ContextToMetadata(newContextForTest(t))
	md, ok := metadata.FromOutgoingContext(outgoingCtx)
	if !ok {
		t.Fatal("expected outgoing metadata")
	}

	incomingCtx := metadata.NewIncomingContext(context.Background(), md)
	ctx := MetadataToContext(incomingCtx)

	assertRequestContextForTest(t, ctx)
}

func TestMetadataToContextRestoresPermissionList(t *testing.T) {
	incomingCtx := metadata.NewIncomingContext(
		context.Background(),
		metadata.Pairs(
			string(typeObject.PermissionListHeaderName),
			typeObject.PermissionUserRead+","+typeObject.PermissionUserWrite,
		),
	)
	ctx := MetadataToContext(incomingCtx)

	requestContext := requestContextMiddleware.GetRequestContext(ctx)
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

func TestContextToMetadataWithoutRequestContextReturnsOriginalContext(t *testing.T) {
	ctx := context.Background()

	got := ContextToMetadata(ctx)

	if got != ctx {
		t.Fatal("expected original context")
	}
	if _, ok := metadata.FromOutgoingContext(got); ok {
		t.Fatal("expected no outgoing metadata")
	}
}

func newContextForTest(t *testing.T) context.Context {
	t.Helper()

	traceID := "123e4567-e89b-12d3-a456-426614174000"
	clientIP := "192.0.2.1"
	userAgent := "go-test"
	userID := "user-1"
	accessToken := "access-token"
	tenantID := "tenant-1"
	locale := "ja-JP"
	timeZone := "AsiaTokyo"

	requestContext, err := requestContextMiddleware.NewRequestContext(
		&requestContextMiddleware.NewRequestContextArgs{
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

	return context.WithValue(
		context.Background(),
		requestContextMiddleware.RequestContextContextName,
		*requestContext,
	)
}

func assertRequestContextForTest(t *testing.T, ctx context.Context) {
	t.Helper()

	requestContext := requestContextMiddleware.GetRequestContext(ctx)
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

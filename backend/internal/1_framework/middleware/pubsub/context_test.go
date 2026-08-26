package pubsub_middleware

import (
	"context"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	requestContextMiddleware "backend/internal/1_framework/middleware/request_context"
	typeObject "backend/internal/4_domain/type_object"
)

func TestContextToHeaderIncludesPermissionList(t *testing.T) {
	requestContext := newRequestContextForTest(t)
	ctx := context.WithValue(
		context.Background(),
		requestContextMiddleware.RequestContextContextName,
		*requestContext,
	)

	headers := ContextToHeader(ctx)
	got := findHeaderValue(headers, string(typeObject.PermissionListHeaderName))

	if got != typeObject.PermissionUserRead+","+typeObject.PermissionUserWrite {
		t.Fatalf("unexpected permissions header: %q", got)
	}
}

func TestHeaderToContextRestoresPermissionList(t *testing.T) {
	ctx := HeaderToContext([]kafka.Header{
		{
			Key: string(typeObject.PermissionListHeaderName),
			Value: []byte(
				typeObject.PermissionUserRead + "," + typeObject.PermissionUserWrite,
			),
		},
	})

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

func TestContextToHeaderWithoutRequestContextReturnsEmptyHeaders(t *testing.T) {
	headers := ContextToHeader(context.Background())

	if len(headers) != 0 {
		t.Fatalf("expected empty headers, got %d", len(headers))
	}
}

func newRequestContextForTest(t *testing.T) *requestContextMiddleware.RequestContext {
	t.Helper()

	requestContext, err := requestContextMiddleware.NewRequestContext(
		&requestContextMiddleware.NewRequestContextArgs{
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

func findHeaderValue(headers []kafka.Header, key string) string {
	for _, header := range headers {
		if header.Key == key {
			return string(header.Value)
		}
	}

	return ""
}

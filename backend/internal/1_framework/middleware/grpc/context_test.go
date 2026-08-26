package grpc_middleware

import (
	"context"
	"testing"

	"google.golang.org/grpc/metadata"

	requestContextMiddleware "backend/internal/1_framework/middleware/request_context"
	typeObject "backend/internal/4_domain/type_object"
)

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

package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	TenantIDMetaName    primitiveObject.ContextKey = "tenant-id"
	TenantIDContextName primitiveObject.ContextKey = "tenantID"
)

var (
	tenantIDMaxLength uint = 99
	tenantIDMinLength uint = 1
)

type TenantID struct {
	err     error
	content *primitiveObject.PrimitiveString
}

func NewTenantID(
	ctx context.Context,
	value *string,
) (
	tenantID TenantID,
) {
	tenantID = TenantID{}
	tenantID.SetValue(ctx, value)

	return
}

func (receiver *TenantID) SetValue(
	ctx context.Context,
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&tenantIDMaxLength),
		primitiveString.WithMinLength(&tenantIDMinLength),
	)

	if receiver.content.GetError() != nil {
		receiver.SetError(
			ctx,
			receiver.content.GetError(),
		)
	}
}
func (receiver *TenantID) GetError() error {
	return receiver.err
}

func (receiver *TenantID) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
}

func (receiver *TenantID) GetValue() string {
	return receiver.content.GetValue()
}

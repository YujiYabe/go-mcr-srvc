package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

const (
	TenantIDHeaderName  primitiveObject.ContextKey = "tenant-id"
	TenantIDContextName primitiveObject.ContextKey = "tenantID"
)

var (
	tenantIDMaxLength uint = 99
	tenantIDMinLength uint = 1
)

type TenantID struct {
	content *primitiveObject.PrimitiveString
}

func NewTenantID(
	value *string,
) (
	tenantID TenantID,
	err error,
) {
	tenantID = TenantID{}
	err = tenantID.setValue(value)

	return
}

func (receiver *TenantID) setValue(
	value *string,
) error {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&tenantIDMaxLength),
		primitiveString.WithMinLength(&tenantIDMinLength),
	)
	if err := receiver.content.Validation(); err != nil {
		return err
	}
	return nil
}

func (receiver TenantID) GetValue() string {
	return receiver.content.GetValue()
}

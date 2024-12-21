package value_object

import (
	"backend/internal/4_domain/primitive_object"
)

type Name struct {
	err             error
	primitiveString *primitive_object.PrimitiveString
}

func NewName() *Name {
	return &Name{
		primitiveString: primitive_object.NewPrimitiveString(),
	}
}

func (object *Name) GetError() error {
	return object.err
}

func (object *Name) SetError(err error) *Name {
	if object.err == nil {
		object.err = err
	}
	return object
}

func (object *Name) GetPointer() *string {
	return object.primitiveString.GetPointer()
}

func (object *Name) GetValue() string {
	return object.primitiveString.GetValue()
}

func (object *Name) SetValue(v *string) *Name {
	p, err := object.primitiveString.SetValue(v)
	if err != nil {
		object.SetError(err)
	}
	object.primitiveString = p

	return object
}

func (object *Name) SetCanNil(v bool) *Name {
	object.primitiveString = object.primitiveString.SetCanNil(v)
	return object
}

func (object *Name) SetMax(v int) *Name {
	object.primitiveString = object.primitiveString.SetMax(v)
	return object
}

func (object *Name) SetMin(v int) *Name {
	object.primitiveString = object.primitiveString.SetMin(v)
	return object
}

func (object *Name) IsNil() bool {
	return object.primitiveString.IsNil()
}

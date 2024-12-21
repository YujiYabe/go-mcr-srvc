package struct_object

import "backend/internal/4_domain/value_object"

type Person struct {
	err  error
	Name value_object.Name
}

type NewPersonArgs struct {
	Name *string
}

func (receiver *Person) GetError() error {
	return receiver.err
}

func (receiver *Person) SetError(err error) *Person {
	if receiver.err == nil {
		receiver.err = err
	}
	return receiver
}

func NewAssociation(args *NewPersonArgs) *Person {

	return &Person{}
}

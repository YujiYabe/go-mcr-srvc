package struct_object

import "backend/internal/4_domain/value_object"

type Person struct {
	err  error
	ID   value_object.ID
	Name value_object.Name
}

type NewPersonArgs struct {
	ID   int
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

func NewPerson(args *NewPersonArgs) (
	person *Person,
) {
	var err error
	person = &Person{}

	person.ID, err = value_object.NewID(args.ID)
	if err != nil {
		person.SetError(err)
		return
	}

	person.Name, err = value_object.NewName(args.Name)
	if err != nil {
		person.SetError(err)
		return
	}

	return
}

package struct_object

import (
	"backend/internal/4_domain/value_object"
)

type PersonList []Person

type Person struct {
	Err         error
	ID          value_object.ID
	Name        value_object.Name
	MailAddress value_object.MailAddress
}

type NewPersonArgs struct {
	ID          *int
	Name        *string
	MailAddress *string
}

func (receiver *Person) GetError() error {
	return receiver.Err
}

func (receiver *Person) SetError(err error) *Person {
	if receiver.Err == nil {
		receiver.Err = err
	}

	return receiver
}

func NewPerson(
	args *NewPersonArgs,
) (
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

	person.MailAddress, err = value_object.NewMailAddress(args.MailAddress)
	if err != nil {
		person.SetError(err)
		return
	}

	return
}

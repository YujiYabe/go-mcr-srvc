package struct_object

import valueObject "backend/internal/4_domain/value_object"

type PersonList []Person

type Person struct {
	Err         error
	ID          valueObject.ID
	Name        valueObject.Name
	MailAddress valueObject.MailAddress
}

type NewPersonArgs struct {
	ID          *int
	Name        *string
	MailAddress *string
}

func (receiver *Person) GetError() error {
	return receiver.Err
}

func (receiver *Person) SetError(
	err error,
) *Person {
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

	person.ID, err = valueObject.NewID(args.ID)
	if err != nil {
		person.SetError(err)
		return
	}

	person.Name, err = valueObject.NewName(args.Name)
	if err != nil {
		person.SetError(err)
		return
	}

	person.MailAddress, err = valueObject.NewMailAddress(args.MailAddress)
	if err != nil {
		person.SetError(err)
		return
	}

	return
}

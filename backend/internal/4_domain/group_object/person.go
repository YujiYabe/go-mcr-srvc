package group_object

import valueObject "backend/internal/4_domain/value_object"

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
) {
	if receiver.Err == nil {
		receiver.Err = err
	}
}

func NewPerson(
	args *NewPersonArgs,
) (
	person *Person,
) {
	person = &Person{}

	person.ID = valueObject.NewID(args.ID)
	if person.ID.Err != nil {
		person.SetError(person.ID.Err)
		return
	}

	person.Name = valueObject.NewName(args.Name)
	if person.Name.Err != nil {
		person.SetError(person.Name.Err)
		return
	}

	person.MailAddress = valueObject.NewMailAddress(args.MailAddress)
	if person.MailAddress.Err != nil {
		person.SetError(person.MailAddress.Err)
		return
	}

	return
}

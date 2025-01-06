package group_object

import valueObject "backend/internal/4_domain/value_object"

type Person struct {
	err         error
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
	return receiver.err
}

func (receiver *Person) SetError(
	err error,
) {
	if receiver.err == nil {
		receiver.err = err
	}
}

func NewPerson(
	args *NewPersonArgs,
) (
	person *Person,
) {
	person = &Person{}

	person.ID = valueObject.NewID(args.ID)
	if person.ID.GetError() != nil {
		person.SetError(person.ID.GetError())
		return
	}

	person.Name = valueObject.NewName(args.Name)
	if person.Name.GetError() != nil {
		person.SetError(person.Name.GetError())
		return
	}

	person.MailAddress = valueObject.NewMailAddress(args.MailAddress)
	if person.MailAddress.GetError() != nil {
		person.SetError(person.MailAddress.GetError())
		return
	}

	return
}

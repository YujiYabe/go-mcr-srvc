package group_object

import typeObject "backend/internal/4_domain/type_object"

type Person struct {
	id          typeObject.ID
	name        typeObject.Name
	mailAddress typeObject.MailAddress
}

type NewPersonArgs struct {
	ID          *int
	Name        *string
	MailAddress *string
}

func NewPerson(
	args *NewPersonArgs,
) (
	person *Person,
	err error,
) {
	person = &Person{}

	person.id, err = typeObject.NewID(args.ID)
	if err != nil {
		return nil, err
	}

	person.name, err = typeObject.NewName(args.Name)
	if err != nil {
		return nil, err
	}

	person.mailAddress, err = typeObject.NewMailAddress(args.MailAddress)
	if err != nil {
		return nil, err
	}

	return
}

func (receiver Person) ID() typeObject.ID {
	return receiver.id
}

func (receiver Person) Name() typeObject.Name {
	return receiver.name
}

func (receiver Person) MailAddress() typeObject.MailAddress {
	return receiver.mailAddress
}

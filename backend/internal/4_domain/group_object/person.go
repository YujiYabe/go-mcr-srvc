package group_object

import (
	"fmt"

	typeObject "backend/internal/4_domain/type_object"
)

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
	if args == nil {
		args = &NewPersonArgs{}
	}

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

func NewPersonSearchCondition(
	args *NewPersonArgs,
) (
	person *Person,
	err error,
) {
	person, err = NewPerson(args)
	if err != nil {
		return nil, err
	}
	if !person.CanBeUsedAsSearchCondition() {
		return nil, fmt.Errorf("person search condition is required")
	}

	return person, nil
}

func ReconstructPerson(
	args *NewPersonArgs,
) (
	person *Person,
	err error,
) {
	person, err = NewPerson(args)
	if err != nil {
		return nil, err
	}
	if !person.HasIdentity() {
		return nil, fmt.Errorf("person identity is required")
	}

	return person, nil
}

func (receiver Person) ID() typeObject.ID {
	return receiver.id
}

func (receiver Person) Identity() typeObject.ID {
	return receiver.id
}

func (receiver Person) Name() typeObject.Name {
	return receiver.name
}

func (receiver Person) MailAddress() typeObject.MailAddress {
	return receiver.mailAddress
}

func (receiver Person) HasIdentity() bool {
	return receiver.id.GetValue() > 0
}

func (receiver Person) HasName() bool {
	return !receiver.name.GetIsNil() && receiver.name.GetValue() != ""
}

func (receiver Person) HasMailAddress() bool {
	return !receiver.mailAddress.GetIsNil() && receiver.mailAddress.GetValue() != ""
}

func (receiver Person) CanBeUsedAsSearchCondition() bool {
	return receiver.HasName() || receiver.HasMailAddress()
}

func (receiver Person) EnsureReadyToUpdate() error {
	if !receiver.HasIdentity() {
		return fmt.Errorf("person identity is required")
	}
	if !receiver.HasName() {
		return fmt.Errorf("person name is required")
	}
	if !receiver.HasMailAddress() {
		return fmt.Errorf("person mail address is required")
	}

	return nil
}

func (receiver *Person) Rename(
	value *string,
) error {
	name, err := typeObject.NewName(value)
	if err != nil {
		return err
	}
	receiver.name = name

	return nil
}

func (receiver *Person) ChangeMailAddress(
	value *string,
) error {
	mailAddress, err := typeObject.NewMailAddress(value)
	if err != nil {
		return err
	}
	receiver.mailAddress = mailAddress

	return nil
}

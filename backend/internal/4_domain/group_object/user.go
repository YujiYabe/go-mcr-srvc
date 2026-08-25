package group_object

import (
	"fmt"

	typeObject "backend/internal/4_domain/type_object"
)

type User struct {
	id    typeObject.ID
	name  typeObject.Name
	email typeObject.Email
}

type NewUserArgs struct {
	ID    *int
	Name  *string
	Email *string
}

func NewUser(
	args *NewUserArgs,
) (
	user *User,
	err error,
) {
	user = &User{}
	if args == nil {
		args = &NewUserArgs{}
	}

	user.id, err = typeObject.NewID(args.ID)
	if err != nil {
		return nil, err
	}

	user.name, err = typeObject.NewName(args.Name)
	if err != nil {
		return nil, err
	}

	user.email, err = typeObject.NewEmail(args.Email)
	if err != nil {
		return nil, err
	}

	return
}

func NewUserSearchCondition(
	args *NewUserArgs,
) (
	user *User,
	err error,
) {
	user, err = NewUser(args)
	if err != nil {
		return nil, err
	}
	if !user.CanBeUsedAsSearchCondition() {
		return nil, fmt.Errorf("user search condition is required")
	}

	return user, nil
}

func ReconstructUser(
	args *NewUserArgs,
) (
	user *User,
	err error,
) {
	user, err = NewUser(args)
	if err != nil {
		return nil, err
	}
	if !user.HasIdentity() {
		return nil, fmt.Errorf("user identity is required")
	}

	return user, nil
}

func (receiver User) ID() typeObject.ID {
	return receiver.id
}

func (receiver User) Identity() typeObject.ID {
	return receiver.id
}

func (receiver User) Name() typeObject.Name {
	return receiver.name
}

func (receiver User) Email() typeObject.Email {
	return receiver.email
}

func (receiver User) HasIdentity() bool {
	return receiver.id.GetValue() > 0
}

func (receiver User) HasName() bool {
	return !receiver.name.GetIsNil() && receiver.name.GetValue() != ""
}

func (receiver User) HasEmail() bool {
	return !receiver.email.GetIsNil() && receiver.email.GetValue() != ""
}

func (receiver User) CanBeUsedAsSearchCondition() bool {
	return receiver.HasName() || receiver.HasEmail()
}

func (receiver User) EnsureReadyToUpdate() error {
	if !receiver.HasIdentity() {
		return fmt.Errorf("user identity is required")
	}
	if !receiver.HasName() {
		return fmt.Errorf("user name is required")
	}
	if !receiver.HasEmail() {
		return fmt.Errorf("user email is required")
	}

	return nil
}

func (receiver *User) Rename(
	value *string,
) error {
	name, err := typeObject.NewName(value)
	if err != nil {
		return err
	}
	receiver.name = name

	return nil
}

func (receiver *User) ChangeEmail(
	value *string,
) error {
	email, err := typeObject.NewEmail(value)
	if err != nil {
		return err
	}
	receiver.email = email

	return nil
}

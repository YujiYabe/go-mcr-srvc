package group_object

import (
	"context"

	domainObject "backend/internal/4_domain/domain_object"
	"backend/internal/logger"
)

type Person struct {
	err         error
	ID          domainObject.ID
	Name        domainObject.Name
	MailAddress domainObject.MailAddress
}

type NewPersonArgs struct {
	ID          *int
	Name        *string
	MailAddress *string
}

func NewPerson(
	ctx context.Context,
	args *NewPersonArgs,
) (
	person *Person,
) {
	person = &Person{}

	person.ID = domainObject.NewID(ctx, args.ID)
	if person.ID.GetError() != nil {
		person.SetError(ctx, person.ID.GetError())
		return
	}

	person.Name = domainObject.NewName(ctx, args.Name)
	if person.Name.GetError() != nil {
		person.SetError(ctx, person.Name.GetError())
		return
	}

	person.MailAddress = domainObject.NewMailAddress(ctx, args.MailAddress)
	if person.MailAddress.GetError() != nil {
		person.SetError(ctx, person.MailAddress.GetError())
		return
	}

	return
}

func (receiver *Person) GetError() error {
	return receiver.err
}

func (receiver *Person) SetError(
	ctx context.Context,
	err error,
) {
	if receiver.err == nil {
		receiver.err = err
		logger.Logging(ctx, receiver.GetError())
	}
}

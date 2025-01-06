package group_object

import (
	"context"

	"backend/pkg"
)

type PersonList struct {
	err     error
	Content []Person
}

type NewPersonListArgs struct {
	Content []NewPersonArgs
}

func (receiver *PersonList) GetError() error {
	return receiver.err
}

func (receiver *PersonList) SetError(
	ctx context.Context,
	err error,
) {
	if receiver.err == nil {
		receiver.err = err
		pkg.Logging(ctx, err)
	}
}

func NewPersonList(
	ctx context.Context,
	args *NewPersonListArgs,
) (
	personList PersonList,
) {
	personList = PersonList{}

	for _, args := range args.Content {
		person := NewPerson(ctx, &args)
		if person.GetError() != nil {
			personList.SetError(ctx, person.GetError())
			break
		}

		personList.Content = append(
			personList.Content,
			*person,
		)
	}

	return
}

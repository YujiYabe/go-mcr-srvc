package group_object

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
	err error,
) {
	if receiver.err == nil {
		receiver.err = err
	}
}

func NewPersonList(
	args *NewPersonListArgs,
) (
	personList PersonList,
) {
	personList = PersonList{}

	for _, args := range args.Content {
		person := NewPerson(&args)
		if person.GetError() != nil {
			personList.SetError(person.GetError())
			break
		}

		personList.Content = append(
			personList.Content,
			*person,
		)
	}

	return
}

package group_object

type PersonList struct {
	content []Person
}

type NewPersonListArgs struct {
	Content []NewPersonArgs
}

func (receiver PersonList) Content() []Person {
	return receiver.content
}

func NewPersonList(
	args *NewPersonListArgs,
) (
	personList PersonList,
	err error,
) {
	personList = PersonList{}
	if args == nil {
		return
	}

	for _, args := range args.Content {
		person, err := NewPerson(&args)
		if err != nil {
			return personList, err
		}

		personList.content = append(
			personList.content,
			*person,
		)
	}

	return
}

func ReconstructPersonList(
	args *NewPersonListArgs,
) (
	personList PersonList,
	err error,
) {
	personList = PersonList{}
	if args == nil {
		return
	}

	for _, args := range args.Content {
		person, err := ReconstructPerson(&args)
		if err != nil {
			return personList, err
		}

		personList.content = append(
			personList.content,
			*person,
		)
	}

	return
}

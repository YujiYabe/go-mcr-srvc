package group_object

type PersonList struct {
	Err     error
	Content []Person
}

type PersonListArgs struct {
	Content []Person
}

func (receiver *PersonList) GetError() error {
	return receiver.Err
}

func (receiver *PersonList) SetError(
	err error,
) {
	if receiver.Err == nil {
		receiver.Err = err
	}
}

func NewPersonList(
	args *NewPersonArgs,
) (
	personList []PersonList,
) {
	personList = []PersonList{}

	// person.ID = valueObject.NewID(args.ID)
	// if person.ID.Err != nil {
	// 	person.SetError(person.ID.Err)
	// 	return
	// }

	// person.Name = valueObject.NewName(args.Name)
	// if person.Name.Err != nil {
	// 	person.SetError(person.Name.Err)
	// 	return
	// }

	// person.MailAddress = valueObject.NewMailAddress(args.MailAddress)
	// if person.MailAddress.Err != nil {
	// 	person.SetError(person.MailAddress.Err)
	// 	return
	// }

	return
}

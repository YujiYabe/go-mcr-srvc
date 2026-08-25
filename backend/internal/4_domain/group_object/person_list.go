package group_object

import (
	"fmt"

	typeObject "backend/internal/4_domain/type_object"
)

type PersonList struct {
	content []Person
}

type NewPersonListArgs struct {
	Content []NewPersonArgs
}

func (receiver PersonList) Content() []Person {
	return receiver.content
}

func (receiver PersonList) IsEmpty() bool {
	return len(receiver.content) == 0
}

func (receiver PersonList) Count() int {
	return len(receiver.content)
}

func (receiver PersonList) ContainsIdentity(
	id typeObject.ID,
) bool {
	if id.GetValue() <= 0 {
		return false
	}

	for _, person := range receiver.content {
		if person.Identity().GetValue() == id.GetValue() {
			return true
		}
	}

	return false
}

func (receiver PersonList) ContainsMailAddress(
	mailAddress typeObject.MailAddress,
) bool {
	if mailAddress.GetIsNil() || mailAddress.GetValue() == "" {
		return false
	}

	for _, person := range receiver.content {
		if person.MailAddress().GetValue() == mailAddress.GetValue() {
			return true
		}
	}

	return false
}

func (receiver *PersonList) Append(
	person Person,
) error {
	if person.HasIdentity() && receiver.ContainsIdentity(person.Identity()) {
		return fmt.Errorf("person identity is duplicated")
	}
	if person.HasMailAddress() && receiver.ContainsMailAddress(person.MailAddress()) {
		return fmt.Errorf("person mail address is duplicated")
	}

	receiver.content = append(receiver.content, person)

	return nil
}

func (receiver PersonList) EnsureNoDuplicateMailAddress() error {
	known := map[string]struct{}{}
	for _, person := range receiver.content {
		if !person.HasMailAddress() {
			continue
		}

		mailAddress := person.MailAddress().GetValue()
		if _, ok := known[mailAddress]; ok {
			return fmt.Errorf("person mail address is duplicated")
		}
		known[mailAddress] = struct{}{}
	}

	return nil
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

		if err := personList.Append(*person); err != nil {
			return personList, err
		}
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

		if err := personList.Append(*person); err != nil {
			return personList, err
		}
	}

	return
}

package group_object

import (
	"fmt"

	typeObject "backend/internal/4_domain/type_object"
)

type UserList struct {
	content []User
}

type NewUserListArgs struct {
	Content []NewUserArgs
}

func (receiver UserList) Content() (
	users []User,
) {
	users = receiver.content

	return
}

func (receiver UserList) IsEmpty() (
	isEmpty bool,
) {
	isEmpty = len(receiver.content) == 0

	return
}

func (receiver UserList) Count() (
	value int,
) {
	value = len(receiver.content)

	return
}

func (receiver UserList) ContainsIdentity(
	id typeObject.ID,
) (
	ok bool,
) {
	if id.GetValue() <= 0 {
		ok = false

		return
	}

	for _, user := range receiver.content {
		if user.Identity().GetValue() == id.GetValue() {
			ok = true

			return
		}
	}

	ok = false

	return
}

func (receiver UserList) ContainsEmail(
	email typeObject.Email,
) (
	ok bool,
) {
	if email.GetIsNil() || email.GetValue() == "" {
		ok = false

		return
	}

	for _, user := range receiver.content {
		if user.Email().GetValue() == email.GetValue() {
			ok = true

			return
		}
	}

	ok = false

	return
}

func (receiver *UserList) Append(
	user User,
) (
	err error,
) {
	if user.HasIdentity() && receiver.ContainsIdentity(user.Identity()) {
		err = fmt.Errorf("user identity is duplicated")

		return
	}
	if user.HasEmail() && receiver.ContainsEmail(user.Email()) {
		err = fmt.Errorf("user email is duplicated")

		return
	}

	receiver.content = append(receiver.content, user)

	err = nil

	return
}

func (receiver UserList) EnsureNoDuplicateEmail() (
	err error,
) {
	known := map[string]struct{}{}
	for _, user := range receiver.content {
		if !user.HasEmail() {
			continue
		}

		email := user.Email().GetValue()
		if _, ok := known[email]; ok {
			err = fmt.Errorf("user email is duplicated")

			return
		}
		known[email] = struct{}{}
	}

	err = nil

	return
}

func NewUserList(
	args *NewUserListArgs,
) (
	userList UserList,
	err error,
) {
	err = nil
	userList = UserList{}
	if args == nil {

		return
	}

	for _, args := range args.Content {
		user, returnedErr := NewUser(&args)
		if returnedErr != nil {
			err = returnedErr

			return
		}

		if returnedErr := userList.Append(*user); returnedErr != nil {
			err = returnedErr

			return
		}
	}

	return
}

func ReconstructUserList(
	args *NewUserListArgs,
) (
	userList UserList,
	err error,
) {
	err = nil
	userList = UserList{}
	if args == nil {

		return
	}

	for _, args := range args.Content {
		user, returnedErr := ReconstructUser(&args)
		if returnedErr != nil {
			err = returnedErr

			return
		}

		if returnedErr := userList.Append(*user); returnedErr != nil {
			err = returnedErr

			return
		}
	}

	return
}

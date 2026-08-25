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

func (receiver UserList) Content() []User {
	return receiver.content
}

func (receiver UserList) IsEmpty() bool {
	return len(receiver.content) == 0
}

func (receiver UserList) Count() int {
	return len(receiver.content)
}

func (receiver UserList) ContainsIdentity(
	id typeObject.ID,
) bool {
	if id.GetValue() <= 0 {
		return false
	}

	for _, user := range receiver.content {
		if user.Identity().GetValue() == id.GetValue() {
			return true
		}
	}

	return false
}

func (receiver UserList) ContainsEmail(
	email typeObject.Email,
) bool {
	if email.GetIsNil() || email.GetValue() == "" {
		return false
	}

	for _, user := range receiver.content {
		if user.Email().GetValue() == email.GetValue() {
			return true
		}
	}

	return false
}

func (receiver *UserList) Append(
	user User,
) error {
	if user.HasIdentity() && receiver.ContainsIdentity(user.Identity()) {
		return fmt.Errorf("user identity is duplicated")
	}
	if user.HasEmail() && receiver.ContainsEmail(user.Email()) {
		return fmt.Errorf("user email is duplicated")
	}

	receiver.content = append(receiver.content, user)

	return nil
}

func (receiver UserList) EnsureNoDuplicateEmail() error {
	known := map[string]struct{}{}
	for _, user := range receiver.content {
		if !user.HasEmail() {
			continue
		}

		email := user.Email().GetValue()
		if _, ok := known[email]; ok {
			return fmt.Errorf("user email is duplicated")
		}
		known[email] = struct{}{}
	}

	return nil
}

func NewUserList(
	args *NewUserListArgs,
) (
	userList UserList,
	err error,
) {
	userList = UserList{}
	if args == nil {
		return
	}

	for _, args := range args.Content {
		user, err := NewUser(&args)
		if err != nil {
			return userList, err
		}

		if err := userList.Append(*user); err != nil {
			return userList, err
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
	userList = UserList{}
	if args == nil {
		return
	}

	for _, args := range args.Content {
		user, err := ReconstructUser(&args)
		if err != nil {
			return userList, err
		}

		if err := userList.Append(*user); err != nil {
			return userList, err
		}
	}

	return
}

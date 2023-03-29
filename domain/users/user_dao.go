package users

import (
	"bookstore_users-api/utils/date"
	"bookstore_users-api/utils/errors"
	"fmt"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.CreatedAt = result.CreatedAt

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %d already registered ", user.Id))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists ", user.Id))
	}

	user.CreatedAt = date.GetNowString()

	userDB[user.Id] = user

	return nil
}

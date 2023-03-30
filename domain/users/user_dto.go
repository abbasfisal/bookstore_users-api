package users

import (
	"bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"created_at"`
}

func (u *User) Validate() *errors.RestErr {

	if u.Email = strings.TrimSpace(strings.ToLower(u.Email)); u.Email == "" {

		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}

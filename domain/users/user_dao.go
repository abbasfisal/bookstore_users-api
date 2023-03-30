package users

import (
	"bookstore_users-api/databases/mysql"
	"bookstore_users-api/utils/date"
	"bookstore_users-api/utils/errors"
	"fmt"
)

const (
	InsertUserQuery = "INSERT INTO users (first_name , last_name, email ,date_created) VALUES (? ,?,?,?);"
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
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {

	stmt, err := mysql.Db.Prepare(InsertUserQuery)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date.GetNowString()

	insertResult, IErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if IErr != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error while saving user : %s", IErr.Error()))
	}
	userId, lErr := insertResult.LastInsertId()
	if lErr != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error while trying to get last inserted ID : %s", err.Error()))
	}
	user.Id = userId

	return nil
}

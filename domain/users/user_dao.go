package users

import (
	"fmt"
	"github.com/Bone1289/bookstore_user-api/datasource/mysql/user_db"
	"github.com/Bone1289/bookstore_user-api/utils/date_utils"
	"github.com/Bone1289/bookstore_user-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	if err := user_db.Client.Ping(); err != nil {
		panic(err)
	}

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
	currentUser := userDB[user.Id]

	if currentUser != nil {
		if currentUser.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.DateCreated = date_utils.GetNowString()

	userDB[user.Id] = user
	return nil
}

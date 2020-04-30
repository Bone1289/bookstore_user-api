package service

import (
	"github.com/Bone1289/bookstore_user-api/domain/users"
	"github.com/Bone1289/bookstore_user-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}

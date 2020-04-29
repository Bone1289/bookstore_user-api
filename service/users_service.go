package service

import (
	"github.com/Bone1289/bookstore_user-api/domain/users"
)

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}

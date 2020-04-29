package app

import (
	"github.com/Bone1289/bookstore_user-api/controllers/ping"
	"github.com/Bone1289/bookstore_user-api/controllers/user"
)

func mapUrls() {
	route.GET("/ping", ping.Ping)

	route.GET("/users/:user_id", user.GetUser)
	route.POST("/users", user.CreateUser)
}

package app

import (
	"github.com/Bone1289/bookstore_user-api/controllers/ping"
	"github.com/Bone1289/bookstore_user-api/controllers/user"
)

func mapUrls() {
	route.GET("/ping", ping.Ping)

	route.POST("/users", user.Create)
	route.GET("/users/:user_id", user.Get)
	route.PUT("/users/:user_id", user.Update)
	route.PATCH("/users/:user_id", user.Update)
	route.DELETE("/users/:user_id", user.Delete)
	route.GET("/internal/users/search", user.Search)
}

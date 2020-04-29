package user

import (
	"encoding/json"
	"fmt"
	"github.com/Bone1289/bookstore_user-api/domain/users"
	"github.com/Bone1289/bookstore_user-api/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		//TODO: Handle error
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		//TODO: Handle error
		return
	}
	result, saveError := service.CreateUser(user)
	if saveError != nil {
		//TODO: handle user creation error
		return
	}

	c.JSON(http.StatusCreated, result)
}

func FindUser(*gin.Context) {}

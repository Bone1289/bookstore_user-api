package user

import (
	"github.com/Bone1289/bookstore_user-api/domain/users"
	"github.com/Bone1289/bookstore_user-api/service"
	"github.com/Bone1289/bookstore_user-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id needs to be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := service.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)

		return
	}

	result, saveError := service.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func FindUser(*gin.Context) {}

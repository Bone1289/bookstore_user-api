package user

import (
	"github.com/Bone1289/bookstore_user-api/domain/users"
	"github.com/Bone1289/bookstore_user-api/service"
	"github.com/Bone1289/bookstore_user-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getUserId(userParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}

	return userId, nil
}

func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := service.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	createdUser, saveError := service.UsersService.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}

	c.JSON(http.StatusCreated, createdUser.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	updatedUser, err := service.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, updatedUser.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := service.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
}

func Search(c *gin.Context) {
	status := c.Query("status")
	userSlice, err := service.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, userSlice.Marshall(c.GetHeader("X-Public") == "true"))
}

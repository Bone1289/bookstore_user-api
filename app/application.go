package app

import (
	"github.com/Bone1289/bookstore_user-api/utils/logger"
	"github.com/gin-gonic/gin"
)

var (
	route = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("about to start the application...")
	route.Run(":8080")
}

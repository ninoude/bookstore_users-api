package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ninoude/bookstore_users-api/logger"
)

var (
	// gin の変数を定義
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("about to start the application...")
	router.Run(":8080")
}

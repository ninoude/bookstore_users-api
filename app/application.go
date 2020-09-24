package app

import (
	"github.com/gin-gonic/gin"
)

var (
	// gin の変数を定義
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	router.Run(":8080")
}

package router

import (
	"github.com/gin-gonic/gin"
	"os"
)

func New() {
	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := router.Run(":" + os.Getenv("APP_PORT"))
	if err != nil {
		panic("App cannot be started")
	}
}

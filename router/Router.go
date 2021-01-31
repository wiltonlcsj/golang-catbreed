package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wiltonlcsj/hgBackendTest/middlewares"
	"github.com/wiltonlcsj/hgBackendTest/requests"
	"os"
)

func New() {
	router := gin.Default()

	router.GET("/ping", requests.Ping)

	router.POST("/login", requests.Login)

	v1 := router.Group("/v1", middlewares.AuthorizeJwt())
	v1.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "authenticated pong",
		})
	})

	err := router.Run(":" + os.Getenv("APP_PORT"))
	if err != nil {
		panic("App cannot be started")
	}
}

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wiltonlcsj/hgBackendTest/middlewares"
	"github.com/wiltonlcsj/hgBackendTest/requests"
	"os"
)

func New() {
	router := gin.Default()

	// Open routes
	router.GET("/ping", requests.Ping)
	router.POST("/login", requests.Login)

	v1 := router.Group("/v1", middlewares.AuthorizeJwt())
	// Must auth routes
	v1.GET("/ping", requests.AuthenticatedPing)
	v1.GET("/breeds", requests.Breed)

	err := router.Run(":" + os.Getenv("APP_PORT"))
	if err != nil {
		panic("App cannot be started")
	}
}

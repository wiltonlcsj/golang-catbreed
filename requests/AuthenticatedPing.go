package requests

import "github.com/gin-gonic/gin"

func AuthenticatedPing(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "authenticated pong",
	})
}

package requests

import "github.com/gin-gonic/gin"

// AuthenticatedPing godoc
// @Security bearerAuth
// @Summary Ping with Authentication
// @Description Test if authenticated ping is working
// @Accept json
// @Produce json
// @Success 200 {object} helpers.HttpDefaultResponse
// @Router /v1/ping [get]
func AuthenticatedPing(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "authenticated pong",
	})
}

package requests

import "github.com/gin-gonic/gin"

// Ping godoc
// @Summary Ping No-Authenticated
// @Description Test if no authenticated ping is working
// @Accept json
// @Produce json
// @Success 200 {object} helpers.HttpDefaultResponse
// @Router /ping [get]
func Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}

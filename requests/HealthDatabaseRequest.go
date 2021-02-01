package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/wiltonlcsj/hgBackendTest/adapters"
)

// HealthDatabase godoc
// @Summary Test database Health
// @Description Test if database ping is working
// @Accept json
// @Produce json
// @Success 200 {object} helpers.HttpDefaultResponse
// @Router /ping [get]
func HealthDatabase(context *gin.Context) {
	dbAdapter := adapters.NewDatabaseAdapter()

	err := dbAdapter.Ping()

	if err != nil {
		context.JSON(200, gin.H{
			"message": "error on ping - "+err.Error(),
		})
		return
	}


	context.JSON(200, gin.H{
		"message": "database is working fine",
	})
}

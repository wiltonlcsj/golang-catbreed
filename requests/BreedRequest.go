package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/wiltonlcsj/hgBackendTest/services"
	"net/http"
)

// Breed godoc
// @Security bearerAuth
// @Summary Search for cat breeds
// @Description Search on API and DB for cat breeds
// @Produce json
// @Accept json
// @Param name query string true "Name"
// @Success 200 {object} []models.CatBreed
// @Failure 400 {object} helpers.HttpDefaultResponse
// @Router /v1/breeds [get]
func Breed(context *gin.Context) {
	var breedService services.SearchBreedService

	if err := context.ShouldBindQuery(&breedService); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	catBreeds, err := services.NewApiSearchService().SearchOnApi(breedService.Name)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusOK, catBreeds)
}

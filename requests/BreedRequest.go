package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/wiltonlcsj/hgBackendTest/services"
	"net/http"
)

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

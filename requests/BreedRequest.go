package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/wiltonlcsj/hgBackendTest/helpers"
	"github.com/wiltonlcsj/hgBackendTest/models"
	"github.com/wiltonlcsj/hgBackendTest/repositories"
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
// @Failure 503 {object} helpers.HttpDefaultResponse
// @Router /v1/breeds [get]
func Breed(context *gin.Context) {
	var breedService services.SearchBreedService
	breedRepository := repositories.NewCatBreedRepository()
	apiRepository := repositories.NewApiSearchRepository()

	if err := context.ShouldBindQuery(&breedService); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var catBreeds []models.CatBreed

	dbConnection := false
	if err := breedRepository.DbAdapter.Ping(); err == nil {
		dbConnection = true
	}

	if dbConnection {
		breeds := breedRepository.FindByName(breedService.Name)
		if len(breeds) != 0 {
			catBreeds = breeds
		}
	}

	apiConnection := true
	if len(catBreeds) == 0 {
		if exists := apiRepository.FindByQueryName(breedService.Name); !exists {
			breeds, err := services.NewApiSearchService().SearchOnApi(breedService.Name)
			apiSearch := &models.ApiSearch{QueryParam: breedService.Name}
			if err != nil {
				apiSearch.ApiStatus = 404
				apiSearch.Response = "Not found"
				apiConnection = false
			} else if len(breeds) != 0 {
				apiSearch.Response = "Not empty"
				catBreeds = breeds
				for i := 0; i < len(catBreeds); i++ {
					existsBreed, err := breedRepository.VerifyIfExists(catBreeds[i].Name)
					if err != nil {
						continue
					}

					if !existsBreed {
						breedRepository.InsertNew(catBreeds[i])
					}
				}
			}

			apiSearch.ApiStatus = 200
			if len(apiSearch.Response) == 0 {
				apiSearch.Response = "empty search"
			}
			apiRepository.InsertNew(*apiSearch)
		}
	}

	if !apiConnection && !dbConnection {
		context.JSON(http.StatusServiceUnavailable, helpers.HttpDefaultResponse{
			Message: "service temporarily unavailable",
		})
		return
	}

	context.JSON(http.StatusOK, catBreeds)
}

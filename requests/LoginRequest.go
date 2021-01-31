package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/wiltonlcsj/hgBackendTest/helpers"
	"github.com/wiltonlcsj/hgBackendTest/models"
	"github.com/wiltonlcsj/hgBackendTest/services"
	"net/http"
)

//A sample user
var testUser = models.User{
	Id:       1,
	Username: "admin",
	Password: "$2a$08$Z1iKDAjpy81dUwILWCbhgeYxioSBneq4pVgzI5.RVS1LjHmZvGcuS",
}

type UserLoginDto struct {
	Username string  `json:"username"`
	Password string `json:"password"`
}

// Login godoc
// @Summary Provides a JSON Web Token
// @Description Authenticates a user and retrieve a JWT for API calls
// @ID Authentication
// @Produce json
// @Accept json
// @Param {object} body requests.UserLoginDto true "Credentials"
// @Success 200 {object} helpers.HttpDefaultResponse
// @Failure 401 {object} helpers.HttpDefaultResponse
// @Failure 400 {object} helpers.HttpDefaultResponse
// @Failure 500 {object} helpers.HttpDefaultResponse
// @Router /login [post]
func Login(context *gin.Context) {
	var userLogin UserLoginDto

	if err := context.ShouldBindJSON(&userLogin); err != nil {
		context.JSON(http.StatusBadRequest, &helpers.HttpDefaultResponse{Message: err.Error()})
		return
	}

	securityService := services.NewSecurityService()
	if err := securityService.ComparePasswords(testUser.Password, userLogin.Password); err != nil {
		context.JSON(http.StatusUnauthorized, &helpers.HttpDefaultResponse{Message: "Invalid login credentials"})
		return
	}

	token, err := services.NewJwtService().CreateJwtToken(testUser.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, &helpers.HttpDefaultResponse{Message: err.Error()})
		return
	}

	context.JSON(http.StatusOK, &helpers.HttpTokenResponse{Token: token})
}

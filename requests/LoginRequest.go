package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/wiltonlcsj/hgBackendTest/models"
	"github.com/wiltonlcsj/hgBackendTest/services"
	"net/http"
)

//A sample use
var testUser = models.User{
	Id:       1,
	Username: "username",
	Password: "password",
}

func Login(context *gin.Context) {
	var userLogin models.User

	if err := context.ShouldBindJSON(&userLogin); err != nil {
		context.JSON(http.StatusBadRequest, "Invalid json sent")
		return
	}

	if userLogin.Username != testUser.Username || userLogin.Password != testUser.Password {
		context.JSON(http.StatusUnauthorized, "Invalid login credentials")
		return
	}

	token, err := services.NewJwtService().CreateJwtToken(testUser.Id)
	if err != nil || token == ""{
		context.JSON(http.StatusInternalServerError, "Internal server error, please repeat the request")
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

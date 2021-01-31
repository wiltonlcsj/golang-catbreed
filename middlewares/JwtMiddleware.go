package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/wiltonlcsj/hgBackendTest/services"
	"net/http"
)

const BearerPrefix = "Bearer "

func AuthorizeJwt() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")

		if len(authHeader) == 0 {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "authorization cannot be empty",
			})
			return
		}

		tokenString := authHeader[len(BearerPrefix):]

		err := services.NewJwtService().VerifyToken(tokenString)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
		}
	}
}

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
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len(BearerPrefix):]

		err := services.NewJwtService().VerifyToken(tokenString)
		if err != nil {
			context.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

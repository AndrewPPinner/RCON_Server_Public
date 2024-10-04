package middleware

import (
	token "RCON_Server/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := token.ValidToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Response": "Unauthorized"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

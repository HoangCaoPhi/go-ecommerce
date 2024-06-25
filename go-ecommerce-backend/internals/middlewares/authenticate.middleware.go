package middlewares

import (
	"github.com/HoangCaoPhi/go-ecommerce/pkg/responses"
	"github.com/gin-gonic/gin"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "valid-token" {
			responses.ErrorReponse(ctx, responses.ErrInvalidToken)
			ctx.Abort()
		}
		ctx.Next()
	}
}

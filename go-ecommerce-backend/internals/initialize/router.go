package initialize

import (
	controller "github.com/HoangCaoPhi/go-ecommerce/internals/controllers"
	"github.com/HoangCaoPhi/go-ecommerce/internals/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.AuthenticateMiddleware())

	v1 := r.Group("/v1")
	{
		v1.GET("/user", controller.NewUserController().GetUserById)
	}

	return r
}

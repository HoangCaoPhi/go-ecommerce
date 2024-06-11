package routers

import (
	controller "github.com/HoangCaoPhi/go-ecommerce/internals/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/user", controller.NewUserController().GetUserById)
	}

	return r
}

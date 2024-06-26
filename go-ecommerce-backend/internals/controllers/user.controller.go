package controller

import (
	"fmt"

	"github.com/HoangCaoPhi/go-ecommerce/internals/services"
	"github.com/HoangCaoPhi/go-ecommerce/pkg/responses"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	fmt.Println("----my handler")
	responses.SuccessResponse(c, 20001, []string{"hcphi test"})
}

package controller

import (
	"best_stock/internal/model"
	"best_stock/internal/service"
	"best_stock/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) Register(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.Error(ctx, 400, "Invalid request")
		return
	}

	if err := c.userService.Register(&user); err != nil {
		response.Error(ctx, 500, "Register failed")
		return
	}

	response.Success(ctx, nil)
}

func (c *UserController) Login(ctx *gin.Context) {
	// 登录逻辑
}

func (c *UserController) GetProfile(context *gin.Context) {
	// 获取个人信息逻辑
}

func (c *UserController) UpdateProfile(context *gin.Context) {
	// 更新个人信息逻辑
}

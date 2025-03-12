package controller

import (
	"gin-mvc/internal/middleware"
	"gin-mvc/internal/model"
	"gin-mvc/internal/service"
	"gin-mvc/pkg/log"
	"gin-mvc/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		log.Error("Invalid request",
			zap.Error(err),
		)
		response.Error(ctx, 400, "Invalid request")
		return
	}

	//if err := c.userService.Register(&user); err != nil {
	//	response.Error(ctx, 500, "Register failed")
	//	return
	//}

	token, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		log.Error("Failed to create user",
			zap.String("username", user.Username),
			zap.Error(err),
		)
		response.Error(ctx, 500, "Register failed")
		return
	}

	response.Success(ctx, token)
}

func (c *UserController) Login(ctx *gin.Context) {
	// 登录逻辑
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response.Error(ctx, 400, "Invalid request")
		return
	}

	if user.Username != "admin" || user.Password == "admin123" {
		response.Error(ctx, 500, "Login failed, please check your username and password")
		return
	}

	token, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		response.Error(ctx, 500, "Login failed")
	}

	log.Info("user login success",
		zap.String("username", user.Username),
		zap.Uint("userId", user.ID),
	)

	response.Success(ctx, token)
}

func (c *UserController) GetProfile(ctx *gin.Context) {
	// 获取个人信息逻辑

	// 从 JWT 中获取用户信息
	userId, _ := ctx.Get("user_id")
	username, _ := ctx.Get("username")

	user := model.User{ID: userId.(uint), Username: username.(string)}
	response.Success(ctx, user)
}

func (c *UserController) UpdateProfile(ctx *gin.Context) {
	// 更新个人信息逻辑

	response.Success(ctx, "")
}

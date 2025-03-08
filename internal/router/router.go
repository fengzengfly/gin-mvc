package router

import (
	"best_stock/internal/controller"
	"best_stock/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(
	userController *controller.UserController,
) *gin.Engine {
	r := gin.Default()

	// 公共路由（不需要鉴权）
	publicGroup := r.Group("/api/public")
	{
		// 登录注册等公共接口
		publicGroup.POST("/login", userController.Login)
		publicGroup.POST("/register", userController.Register)

		// 可选鉴权的接口
		//publicGroup.GET("/products", middleware.OptionalJWTMiddleware(), productController.ListProducts)
	}

	// 需要鉴权的路由
	authGroup := r.Group("/api")
	authGroup.Use(middleware.JWTMiddleware())
	{
		// 用户相关
		authGroup.GET("/user/profile", userController.GetProfile)
		authGroup.PUT("/user/update", userController.UpdateProfile)

	}

	return r
}

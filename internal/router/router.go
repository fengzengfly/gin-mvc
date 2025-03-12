package router

import (
	"gin-mvc/internal/middleware"
	"gin-mvc/internal/modules"
	"github.com/gin-gonic/gin"
)

func InitRouter(modules *modules.Modules) *gin.Engine {

	r := gin.Default()

	// 公共路由（不需要鉴权）
	publicGroup := r.Group("/api/public")
	publicGroup.Use(middleware.RequestLogger(), middleware.ErrorHandler())
	{
		// 登录注册等公共接口
		publicGroup.POST("/user/login", modules.UserController.Login)
		publicGroup.POST("/user/register", modules.UserController.Register)

		// 可选鉴权的接口
		//publicGroup.GET("/products", middleware.OptionalJWTMiddleware(), productController.ListProducts)
	}

	// 需要鉴权的路由
	authGroup := r.Group("/api")
	authGroup.Use(middleware.RequestLogger(), middleware.ErrorHandler(), middleware.JWTMiddleware())
	{
		// 用户相关
		authGroup.GET("/user/profile", modules.UserController.GetProfile)
		authGroup.PUT("/user/update", modules.UserController.UpdateProfile)

	}
	return r
}

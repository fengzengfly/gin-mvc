package modules

import (
	"gin-mvc/internal/controller"
	"gin-mvc/internal/repository"
	"gin-mvc/internal/service"
	"gorm.io/gorm"
)

type Modules struct {
	UserController *controller.UserController
	// 其他模块控制器
}

func InitModules(db *gorm.DB) *Modules {
	// 用户模块
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	return &Modules{
		UserController: userController,
	}
}

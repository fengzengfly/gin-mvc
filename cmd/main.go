package main

import (
	"best_stock/internal/controller"
	"best_stock/internal/repository"
	"best_stock/internal/router"
	"best_stock/internal/service"
	"best_stock/pkg/config"
	"best_stock/pkg/database"
	"fmt"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// 初始化数据库
	db, err := database.InitDatabase(&cfg.Database)
	if err != nil {
		panic(err)
	}

	// 依赖注入
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	// 初始化路由
	r := router.InitRouter(userController)

	// 启动服务
	err = r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"gin-mvc/internal/modules"
	"gin-mvc/internal/router"
	"gin-mvc/pkg/config"
	"gin-mvc/pkg/database"
	"gin-mvc/pkg/log"
	"go.uber.org/zap"
)

func main() {
	// 加载系统配置
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// 初始化日志
	if err := log.InitLogger(&cfg.Log); err != nil {
		panic(err)
	}

	// 初始化数据库
	db, err := database.InitDatabase(&cfg.Database)
	if err != nil {
		panic(err)
	}

	// 初始化模块
	initedModules := modules.InitModules(db)

	// 初始化路由
	r := router.InitRouter(initedModules)

	// 启动服务
	log.Info("Server starting",
		zap.String("port", fmt.Sprintf(":%d", cfg.Server.Port)),
	)

	if err := r.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		log.Fatal("Server start failed", zap.Error(err))
	}
}

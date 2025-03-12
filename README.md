# GIN-MVC - Gin MVC Framework

A Go web application built using the Gin framework with an MVC architecture inspired by Java Spring MVC patterns.

## 项目概述

GIN-MVC 项目采用类似 Spring MVC 的分层架构设计，使用 Go 语言和 Gin 框架实现。项目结构清晰，实现了控制层、服务层和数据访问层的分离，并使用依赖注入的方式组织各组件间的关系。

## 核心特性

- **分层架构**: 遵循 MVC 模式，清晰分离控制器、服务和数据访问层
- **依赖注入**: 通过构造函数注入依赖，降低组件间耦合
- **模块化设计**: 可扩展的模块化结构，便于添加新功能
- **配置管理**: 外部化配置，支持不同环境
- **日志系统**: 集成 zap 日志框架

## 项目结构

```
/
├── cmd/
│   └── main.go              # 应用程序入口
├── internal/
│   ├── controller/          # 控制器层，处理HTTP请求
│   ├── model/               # 数据模型定义
│   ├── modules/             # 模块初始化与管理
│   ├── repository/          # 数据访问层，处理数据持久化
│   ├── router/              # 路由配置
│   └── service/             # 业务逻辑层
└── pkg/
    ├── config/              # 配置管理
    ├── database/            # 数据库连接
    └── log/                 # 日志工具
```

## Spring MVC VS Gin MVC 对比

| Spring MVC | Gin MVC |
|------------|---------|
| Controller | Controller |
| Service | Service |
| Repository | Repository |
| Dependency Injection | Constructor Injection |
| Bean Configuration | Modules Setup |
| Properties Files | Viper Configuration |

## 快速开始

### 前置要求

- Go 1.23+
- MySQL

### 安装与运行

```bash
# 克隆项目
git clone https://github.com/fengzengfly/gin-mvc.git

# 安装依赖
go mod tidy

# 运行应用
go run cmd/main.go
```

## 开发指南

### 添加新模块

1. 在 `repository` 中创建数据访问层
2. 在 `service` 中实现业务逻辑
3. 在 `controller` 中添加API接口
4. 在 `modules/module.go` 中注册新模块
5. 在 `router` 中添加路由

## 使用的技术

- Gin: Web框架
- GORM: ORM库
- Zap: 日志框架
- Viper: 配置管理

## 架构设计理念

本项目借鉴了Java Spring MVC的设计思想，通过清晰的分层和依赖注入实现了关注点分离和代码复用，同时利用Go语言的简洁和高效特性，打造轻量级但功能完备的Web应用框架。

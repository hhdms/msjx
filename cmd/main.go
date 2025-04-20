package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"main/internal/api/v1"
	"main/internal/app"
)

func main() {
	// 加载配置
	app.LoadConfig()

	// 初始化数据库连接
	app.InitDB()

	// 创建Gin引擎
	r := gin.Default()

	// 注册路由
	v1.RegisterRoutes(r)

	// 启动服务器
	log.Println("服务器启动在 :8080 端口")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
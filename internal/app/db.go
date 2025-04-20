package app

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"main/internal/models"
)

// DB 全局数据库连接
var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	// 加载配置文件
	LoadConfig()

	// 从配置获取数据库连接信息
	host := AppConfig.Database.Host
	port := AppConfig.Database.Port
	username := AppConfig.Database.User
	password := AppConfig.Database.Password
	dbname := AppConfig.Database.Name

	// 构建DSN连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	// 配置GORM
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// 连接数据库
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 设置连接池
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxConn)

	// 自动迁移表结构
	err = DB.AutoMigrate(&models.Dept{})
	if err != nil {
		log.Fatalf("自动迁移表结构失败: %v", err)
	}

	log.Println("数据库连接成功")
}

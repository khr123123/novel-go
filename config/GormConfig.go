package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB(user, password, host string, port int, dbname string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 设置连接池参数（可选）
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("获取数据库连接对象失败: %v", err)
	}
	// 最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	// 最大打开连接数
	sqlDB.SetMaxOpenConns(100)
	// 连接最大生命周期
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("数据库连接成功")
}

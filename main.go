package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"novel-go/config"
	"novel-go/controller"
	_ "novel-go/docs"
	"novel-go/service"
)

func main() {
	//  1. connect mysql databases
	config.InitDB("root", "admin123", "127.0.0.1", 3306, "novel")
	//   connect redis
	pong, err := config.RedisClient.Ping(config.Ctx).Result()
	if err != nil || pong != "PONG" {
		log.Fatalf("❌ Redis连接失败: %v", err)
	} else {
		log.Println("✅ Redis连接成功")
	}
	//  2.  run the gin frame
	r := gin.Default()
	//  3.  cors config
	r.Use(config.Cors())
	//  4. 1注册swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//  5. 1 Resource Related APIs
	frontApiGroup := r.Group("/api/front")
	resourceController := controller.NewResourceController(service.NewResourceServiceImpl()) //Dependency Inversion
	resourceController.RegisterRoutes(frontApiGroup)
	//  5. 2 User Related APIs
	userController := controller.NewUserController(service.NewUserServiceImpl()) //Dependency Inversion
	userController.RegisterRoutes(frontApiGroup)
	//  5. 3 Home Related APIs
	homeController := controller.NewHomeController(service.NewHomeService()) //Dependency Inversion
	homeController.RegisterRoutes(frontApiGroup)

	_ = r.Run(":8888")

}

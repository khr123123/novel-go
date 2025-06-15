package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"novel-go/config"
	"novel-go/controller"
	_ "novel-go/docs"
	"novel-go/service"
)

func main() {
	//  1. connect mysql databases
	config.InitDB("root", "admin123", "127.0.0.1", 3306, "novel")
	//  2.  run the gin frame
	r := gin.Default()
	//  2. 1注册swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//  2. 2 Resource Related APIs
	frontApiGroup := r.Group("/api/front")
	resourceController := controller.NewResourceController(service.NewResourceService()) //Dependency Inversion
	resourceController.RegisterRoutes(frontApiGroup)
	//  2. 3 User Related APIs
	userController := controller.NewUserController(service.NewUserService()) //Dependency Inversion
	userController.RegisterRoutes(frontApiGroup)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}

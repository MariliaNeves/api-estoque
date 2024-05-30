package main

import (
	"github.com/MariliaNeves/api-estoque/src/configuration/logger"
	"github.com/MariliaNeves/api-estoque/src/controller"
	"github.com/MariliaNeves/api-estoque/src/controller/routes"
	"github.com/MariliaNeves/api-estoque/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Start App")
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
	}

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		logger.Error("Error", err)
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MariliaNeves/api-estoque/src/configuration/logger"
	"github.com/MariliaNeves/api-estoque/src/controler/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Start App")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("TEST"))

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

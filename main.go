package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lucasquitan/users-api-go/src/configuration/logger"
	"github.com/lucasquitan/users-api-go/src/controller/routes"
)

func main() {
	logger.Info("About to start user appication")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Error to expose the application", err)
	}
}

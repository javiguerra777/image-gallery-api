package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"image-gallery-server/services"
	"image-gallery-server/database"
	"image-gallery-server/endpoint"
)

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Server is running",
	})
}
func main() {
	var port string
	if os.Getenv("ENV") == "development" {
		fmt.Println("Loading .env file")
		err := services.LoadGodotEnv()
		if err != nil {
			os.Exit(1)
		}
		port = "127.0.0.1:4000"
	} else {
		port = ":4000"
	}
	cfg := services.LoadConfig()
	os.Setenv(("AWS_ACCESS_KEY_ID"), cfg.AwsConfig.AWS_ACCESS_KEY_ID)
	os.Setenv(("AWS_SECRET_ACCESS_KEY"), cfg.AwsConfig.AWS_SECRET_ACCESS_KEY)
	dynamodbClient, err := database.ConnectToDynamoDB(&cfg.AwsConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	postgresdbClient, err := database.ConnectToPostgresDB(&cfg.PostgresConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	router := gin.Default()
	// public routes
	router.GET("/health", HealthCheck)
	// private routes
	endpoint.ImageHandler(router, dynamodbClient, postgresdbClient)
	endpoint.GalleriesHandler(router, dynamodbClient, postgresdbClient)
	fmt.Println("Server is running")
	router.Run(port)
}
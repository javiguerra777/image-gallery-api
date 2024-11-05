package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"image-gallery-server/services"
	"image-gallery-server/database"
	"image-gallery-server/endpoint"
)

func main() {
	if os.Getenv("ENV") == "development" {
		fmt.Println("Loading .env file")
		err := services.LoadGodotEnv()
		if err != nil {
			os.Exit(1)
		}
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
	endpoint.ImageHandler(router, dynamodbClient, postgresdbClient)
	endpoint.GalleriesHandler(router, dynamodbClient, postgresdbClient)
	fmt.Println("Server is running")
	router.Run(os.Getenv("PORT"))
}
package services

import (
	"os"
	"image-gallery-server/models"
)

func LoadConfig() *models.Config {
	cfg := &models.Config{
		AwsConfig: models.AwsConfig{
			Region: os.Getenv("AWS_REGION"),
			AWS_SECRET_ACCESS_KEY: os.Getenv("AWS_SECRET_ACCESS_KEY"),
			AWS_ACCESS_KEY_ID: os.Getenv("AWS_ACCESS_KEY_ID"),
		},
		PostgresConfig: models.PostgresConfig{
			DB_NAME: os.Getenv("DB_NAME"),
			DB_USER: os.Getenv("DB_USER"),
			DB_PASSWORD: os.Getenv("DB_PASSWORD"),
			DB_PORT: os.Getenv("DB_PORT"),
			DB_HOST: os.Getenv("DB_HOST"),
			SSL_MODE: os.Getenv("SSL_MODE"),
		},
	}
	return cfg
}
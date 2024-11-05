package services

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadGodotEnv() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return err
	}
	return nil
}
package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"image-gallery-server/models"
)
func ConnectToPostgresDB(cfg *models.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_PORT, cfg.SSL_MODE)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
			return nil, err
	}
	err = db.AutoMigrate(&models.Gallery{}, &models.Image{})
	if err != nil {
			return nil, err
	}

	fmt.Println("Successfully connected to PostgreSQL using GORM!")
	return db, nil
}
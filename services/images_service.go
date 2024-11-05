package services

import (
	"image-gallery-server/models"

	"gorm.io/gorm"
)

func FindImages(postgresdbClient *gorm.DB) (img []models.Image, err error) {
	var images []models.Image
	if err := postgresdbClient.Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

func FindImageById(postgresdbClient *gorm.DB, imageId uint) (img models.Image, err error) {
  var image models.Image
	if err := postgresdbClient.First(&image, imageId).Error; err != nil {
		return models.Image{}, err
	}
	return image, nil
}

func CreateImage(postgresdbClient *gorm.DB, image models.Image) error {
if err := postgresdbClient.Create(&image).Error; err != nil {
	return err
}
	return nil
}

func DeleteImage(postgresdbClient *gorm.DB) {}
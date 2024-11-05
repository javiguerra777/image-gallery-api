package services

import (
	"gorm.io/gorm"
	"image-gallery-server/models"
)

func FindGalleries(postgresdbClient *gorm.DB) (galls []models.Gallery, err error) {
	var galleries []models.Gallery
	if err := postgresdbClient.Find(&galleries).Error; err != nil {
		return nil, err
	}
	return galleries, nil
}

func FindGalleryByID(postgresdbClient *gorm.DB, galleryId uint) (gall models.Gallery, err error) {
	var gallery models.Gallery
	if err := postgresdbClient.First(&gallery, galleryId).Error; err != nil {
		return models.Gallery{}, err
	}
	return gallery, nil
}

func CreateGallery(postgresdbClient *gorm.DB, gallery models.Gallery) error {
	if err := postgresdbClient.Create(&gallery).Error; err != nil {
		return err
	}
	return nil
}

func DeleteGallery() {}
package models

import "gorm.io/gorm"

type Gallery struct {
	gorm.Model
	Title 		 string `json:"title"`
	Description  string `json:"description"`
	UserID 		 uint `json:"user_id"`
	CoverImageUrl string `json:"cover_image_url"`
	Images []Image `json:"images" gorm:"foreignKey:GalleryID"`
}
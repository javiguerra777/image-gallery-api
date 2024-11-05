package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ImageURL string `json:"image_url"`
	Title		string `json:"title"`
	Description string `json:"description"`
	UserID		uint `json:"user_id" gorm:"not null"`
	GalleryID	uint `json:"gallery_id" gorm:"not null"`
}
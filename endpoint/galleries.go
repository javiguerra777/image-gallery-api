package endpoint

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"image-gallery-server/auth"
	"image-gallery-server/models"
	"image-gallery-server/services"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"gorm.io/gorm"
)

type GalleryForm struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
func GalleriesHandler(router *gin.Engine, dynamodbClient *dynamodb.DynamoDB, postgresdbClient *gorm.DB) {
	router.Use(auth.AuthMiddleware(dynamodbClient))
	router.GET("/api/galleries", func(c *gin.Context) {
		galleries, err := services.FindGalleries(postgresdbClient)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, galleries)
	})
	router.GET("/api/galleries/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
				return
		}
		gallery, err := services.FindGalleryByID(postgresdbClient, uint(id))
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gallery)
	})
	router.POST("/api/galleries", func(c *gin.Context) {
		tokenUser, exists := c.Get("tokenUser")
		if !exists {
			c.JSON(400, gin.H{"error": "User not found"})
			return
		}
		fileHeader, file, err := services.ExtractFile(c, "image")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		var galleryForm GalleryForm
		if err := c.ShouldBind(&galleryForm); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		galleryImageUrl, err := services.UploadFileToS3(file, fileHeader,"image-gallery-bucket-test", "cover-images")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		gallery := models.Gallery {
			Title: galleryForm.Title,
			Description: galleryForm.Description,
			CoverImageUrl: galleryImageUrl,
			UserID: uint(tokenUser.(models.DynamoDBAuthToken).UserId),
		}
		if err := postgresdbClient.Create(&gallery).Error; err != nil {
			if err := services.DeleteFileFromS3("image-gallery-bucket-test", galleryImageUrl); err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Gallery created", "gallery": gallery })
	})
	router.PUT("/api/galleries/:id", func(c *gin.Context) {})
	router.DELETE("/api/galleries/:id", func(c *gin.Context) {})
}
package endpoint

import (
	"image-gallery-server/auth"
	"image-gallery-server/models"
	"image-gallery-server/services"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ImageForm struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	GalleryID uint `json:"gallery_id" binding:"required"`
}
func ImageHandler(router *gin.Engine, dynamodbClient *dynamodb.DynamoDB, postgresdbClient *gorm.DB) {
	router.Use(auth.AuthMiddleware(dynamodbClient))
	router.GET("/api/images", func(c *gin.Context) {
		images, err := services.FindImages(postgresdbClient)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, images)
	})
	router.GET("/api/images/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}
		image, err := services.FindImageById(postgresdbClient, uint(id))
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, image)
	})
	router.POST("/api/images", func(c *gin.Context) {
		tokenUser, exists := c.Get("tokenUser")
		if !exists {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		fileHeader, file, err := services.ExtractFile(c, "image")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		var imageForm ImageForm
		if err := c.ShouldBind(&imageForm); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if _, err := services.FindGalleryByID(postgresdbClient, imageForm.GalleryID); err != nil {
			c.JSON(404, gin.H{"error": "Gallery not found"})
			return
		}
		imageUrl, err := services.UploadFileToS3(file, fileHeader, "image-gallery-bucket-test", "images")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		image := models.Image {
			Title: imageForm.Title,
			Description: imageForm.Description,
			UserID: uint(tokenUser.(models.DynamoDBAuthToken).UserId),
			GalleryID: imageForm.GalleryID,
			ImageURL: imageUrl,
		}
		if err := services.CreateImage(postgresdbClient, image); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "Image Created", "image": image})
	})
	router.PUT("/api/images/:id", func(c *gin.Context) {})
	router.DELETE("/api/images/:id", func(c *gin.Context) {})
}
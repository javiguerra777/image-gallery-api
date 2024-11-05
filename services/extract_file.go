package services

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)


func ExtractFile(c *gin.Context, fileName string) (*multipart.FileHeader, multipart.File, error) {
	fileHeader, err := c.FormFile(fileName)
	if err != nil {
		return nil, nil, err
	}
	file, err := fileHeader.Open()
	if err != nil {
		return nil, nil, err
	}
	return fileHeader, file, nil
}
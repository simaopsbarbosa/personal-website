package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadImageHandler handles image uploads
// @Summary upload image
// @Schemes
// @Description upload an image to the server
// @Tags images
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Image file"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /upload [post]
func UploadImageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no file uploaded"})
			return
		}

		uploadDir := "uploads"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			if err := os.Mkdir(uploadDir, 0755); err != nil {
				log.Println("failed to create upload directory:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save image"})
				return
			}
		}

		// generate a unique filename and validate extension
		ext := filepath.Ext(file.Filename)
		if !isAllowedExtension(ext) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file type not allowed"})
			return
		}

		filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		filePath := filepath.Join(uploadDir, filename)

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			log.Println("failed to save file:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save image"})
			return
		}

		// return the URL
		scheme := "http"
		if c.Request.TLS != nil {
			scheme = "https"
		}
		
		host := c.Request.Host
		fileURL := fmt.Sprintf("%s://%s/uploads/%s", scheme, host, filename)

		c.JSON(http.StatusCreated, gin.H{
			"url":      fileURL,
			"filename": filename,
		})
	}
}

func isAllowedExtension(ext string) bool {
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp", ".svg":
		return true
	default:
		return false
	}
}

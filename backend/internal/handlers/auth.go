package handlers

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"simaopsbarbosa/backend/internal/dto"
	"simaopsbarbosa/backend/internal/utils"
)

// LoginHandler handles admin login with a single password check
// @Summary login admin
// @Schemes
// @Description login with admin password
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body dto.UserLoginRequest true "Login payload"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload dto.UserLoginRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		payload.Password = strings.TrimSpace(payload.Password)
		if payload.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
			return
		}

		storedHash := os.Getenv("ADMIN_PASSWORD_HASH")
		if storedHash == "" {
			log.Println("ERROR: ADMIN_PASSWORD_HASH environment variable is not set")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "server configuration error"})
			return
		}

		if !PasswordMatches(storedHash, payload.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		token, err := utils.GenerateToken(24 * time.Hour)
		if err != nil {
			log.Println("failed to generate token:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to login"})
			return
		}

		response := dto.LoginResponse{
			Message: "login successful",
			Token:   token,
		}
		c.JSON(http.StatusOK, response)
	}
}

func PasswordMatches(storedPasswordHash string, providedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(providedPassword)) == nil
}

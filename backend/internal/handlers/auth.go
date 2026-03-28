package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"simaopsbarbosa/backend/internal/dto"
	"simaopsbarbosa/backend/internal/models"
)

// RegisterHandler godoc
// @Summary register user
// @Schemes
// @Description create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body dto.UserRegisterRequest true "Register payload"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func RegisterHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload dto.UserRegisterRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		payload.Name = strings.TrimSpace(payload.Name)
		payload.Email = strings.TrimSpace(payload.Email)
		payload.Password = strings.TrimSpace(payload.Password)

		hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("failed to hash password:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
			return
		}

		result, err := db.Exec(
			"INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
			payload.Name,
			payload.Email,
			string(hash),
		)
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
				return
			}
			log.Println("failed to register user:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
			return
		}

		id, _ := result.LastInsertId()
		response := dto.UserResponse{
			ID:    int(id),
			Name:  payload.Name,
			Email: payload.Email,
		}
		c.JSON(http.StatusCreated, response)
	}
}

// LoginHandler godoc
// @Summary login user
// @Schemes
// @Description login with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body dto.UserLoginRequest true "Login payload"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func LoginHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload dto.UserLoginRequest
		decoder := json.NewDecoder(c.Request.Body)
		if err := decoder.Decode(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		payload.Email = strings.TrimSpace(payload.Email)
		payload.Password = strings.TrimSpace(payload.Password)
		if payload.Email == "" || payload.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
			return
		}

		var user models.User
		err := db.QueryRow(
			"SELECT COALESCE(id, rowid), COALESCE(name, ''), email, password, COALESCE(created_at, ''), COALESCE(updated_at, '') FROM users WHERE email = ?",
			payload.Email,
		).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
				return
			}
			log.Println("failed to login user:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to login"})
			return
		}

		if !PasswordMatches(user.Password, payload.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		response := dto.LoginResponse{
			Message: "login successful",
			User: dto.UserResponse{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
		}
		c.JSON(http.StatusOK, response)
	}
}

func PasswordMatches(storedPassword string, providedPassword string) bool {
	if bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword)) == nil {
		return true
	}

	// Backward compatibility for records that may still store plain text passwords.
	return storedPassword == providedPassword
}

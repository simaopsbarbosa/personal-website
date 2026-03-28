package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	router.POST("/auth/register", registerHandler(db))
	router.POST("/auth/login", loginHandler(db))

	router.POST("/posts", createPostHandler(db))
	router.GET("/posts", getAllPostsHandler(db))
	router.GET("/posts/:slug", getPostHandler(db))
	router.PUT("/posts/:slug", updatePostHandler(db))
	router.DELETE("/posts/:slug", deletePostHandler(db))
}

// RegisterExample godoc
// @Summary register user
// @Schemes
// @Description create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body User true "Register payload"
// @Success 201 {object} User
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func registerHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload User
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
		createdUser := User{ID: int(id), Name: payload.Name, Email: payload.Email}
		c.JSON(http.StatusCreated, createdUser)
	}
}

// LoginExample godoc
// @Summary login user
// @Schemes
// @Description login with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param payload body User true "Login payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func loginHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload User
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

		var user User
		err := db.QueryRow(
			"SELECT COALESCE(id, rowid), COALESCE(name, ''), email, password, COALESCE(created_at, '') FROM users WHERE email = ?",
			payload.Email,
		).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
				return
			}
			log.Println("failed to login user:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to login"})
			return
		}

		if !passwordMatches(user.Password, payload.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		user.Password = ""
		c.JSON(http.StatusOK, gin.H{
			"message": "login successful",
			"user":    user,
		})
	}
}

func passwordMatches(storedPassword string, providedPassword string) bool {
	if bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword)) == nil {
		return true
	}

	// Backward compatibility for records that may still store plain text passwords.
	return storedPassword == providedPassword
}

// CreatePostExample godoc
// @Summary create post
// @Schemes
// @Description create a new blog post
// @Tags posts
// @Accept json
// @Produce json
// @Param payload body Post true "Post payload"
// @Success 201 {object} Post
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts [post]
func createPostHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload Post
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		payload.Title = strings.TrimSpace(payload.Title)
		payload.Content = strings.TrimSpace(payload.Content)
		payload.Slug = strings.TrimSpace(payload.Slug)

		slugSource := payload.Slug
		if slugSource == "" {
			slugSource = payload.Title
		}

		uniqueSlug, err := generateUniquePostSlug(db, slugSource, 0)
		if err != nil {
			log.Println("failed to generate slug:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create post"})
			return
		}
		payload.Slug = uniqueSlug

		result, err := db.Exec(
			"INSERT INTO posts (user_id, title, content, slug) VALUES (?, ?, ?, ?)",
			payload.UserID,
			payload.Title,
			payload.Content,
			payload.Slug,
		)
		if err != nil {
			log.Println("failed to create post:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create post"})
			return
		}

		id, _ := result.LastInsertId()
		createdPost := Post{ID: int(id), UserID: payload.UserID, Slug: payload.Slug, Title: payload.Title, Content: payload.Content}
		c.JSON(http.StatusCreated, createdPost)
	}
}

// GetPostExample godoc
// @Summary get post by slug
// @Schemes
// @Description fetch a post using its slug
// @Tags posts
// @Accept json
// @Produce json
// @Param slug path string true "Post slug"
// @Success 200 {object} Post
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts/{slug} [get]
func getPostHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug, ok := parseSlugParam(c)
		if !ok {
			return
		}

		var post Post

		err := db.QueryRow(
			"SELECT COALESCE(id, rowid), user_id, slug, title, content, COALESCE(created_at, '') FROM posts WHERE slug = ?",
			slug,
		).Scan(&post.ID, &post.UserID, &post.Slug, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
				return
			}
			log.Println("failed to fetch post:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch post"})
			return
		}

		c.JSON(http.StatusOK, post)
	}
}

// GetAllPostsExample godoc
// @Summary get all posts
// @Schemes
// @Description fetch all blog posts
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {object} []Post
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts [get]
func getAllPostsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		posts := []Post{}

		rows, err := db.Query(
			"SELECT COALESCE(id, rowid), user_id, slug, title, content, COALESCE(created_at, '') FROM posts",
		)
		if err != nil {
			log.Println("failed to fetch posts:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch posts"})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var post Post
			if err := rows.Scan(&post.ID, &post.UserID, &post.Slug, &post.Title, &post.Content, &post.CreatedAt); err != nil {
				log.Println("failed to scan post:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch posts"})
				return
			}
			posts = append(posts, post)
		}

		if err := rows.Err(); err != nil {
			log.Println("error iterating posts:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch posts"})
			return
		}

		if len(posts) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "no posts found"})
			return
		}

		c.JSON(http.StatusOK, posts)
	}
}

// UpdatePostExample godoc
// @Summary update post by slug
// @Schemes
// @Description update an existing post using its slug
// @Tags posts
// @Accept json
// @Produce json
// @Param slug path string true "Post slug"
// @Param payload body Post true "Post payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts/{slug} [put]
func updatePostHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug, ok := parseSlugParam(c)
		if !ok {
			return
		}

		var existingID int
		err := db.QueryRow("SELECT COALESCE(id, rowid) FROM posts WHERE slug = ?", slug).Scan(&existingID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
				return
			}
			log.Println("failed to fetch post before update:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update post"})
			return
		}

		var payload Post
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		payload.Title = strings.TrimSpace(payload.Title)
		payload.Content = strings.TrimSpace(payload.Content)
		payload.Slug = strings.TrimSpace(payload.Slug)

		slugSource := payload.Slug
		if slugSource == "" {
			slugSource = payload.Title
		}

		uniqueSlug, err := generateUniquePostSlug(db, slugSource, existingID)
		if err != nil {
			log.Println("failed to generate slug:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update post"})
			return
		}
		payload.Slug = uniqueSlug

		result, err := db.Exec(
			"UPDATE posts SET user_id = ?, title = ?, content = ?, slug = ? WHERE id = ?",
			payload.UserID,
			payload.Title,
			payload.Content,
			payload.Slug,
			existingID,
		)
		if err != nil {
			log.Println("failed to update post:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update post"})
			return
		}

		rows, _ := result.RowsAffected()
		if rows == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "post updated", "slug": payload.Slug})
	}
}

// DeletePostExample godoc
// @Summary delete post by slug
// @Schemes
// @Description delete an existing post using its slug
// @Tags posts
// @Accept json
// @Produce json
// @Param slug path string true "Post slug"
// @Success 204 {string} string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts/{slug} [delete]
func deletePostHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug, ok := parseSlugParam(c)
		if !ok {
			return
		}

		result, err := db.Exec("DELETE FROM posts WHERE slug = ?", slug)
		if err != nil {
			log.Println("failed to delete post:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete post"})
			return
		}

		rows, _ := result.RowsAffected()
		if rows == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func parseIDParam(c *gin.Context) (int, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return 0, false
	}

	return id, true
}

func parseSlugParam(c *gin.Context) (string, bool) {
	slug := strings.TrimSpace(c.Param("slug"))
	if slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid slug"})
		return "", false
	}

	return slug, true
}

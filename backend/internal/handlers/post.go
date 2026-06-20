package handlers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"simaopsbarbosa/backend/internal/dto"
	"simaopsbarbosa/backend/internal/utils"
)

// CreatePostHandler godoc
// @Summary create post
// @Schemes
// @Description create a new blog post
// @Tags posts
// @Accept json
// @Produce json
// @Param payload body dto.PostCreateRequest true "Post payload"
// @Success 201 {object} dto.PostResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts [post]
func CreatePostHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload dto.PostCreateRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		payload.Title = strings.TrimSpace(payload.Title)
		payload.Content = strings.TrimSpace(payload.Content)

		uniqueSlug, err := utils.GenerateUniquePostSlug(db, payload.Title, 0)
		if err != nil {
			log.Println("failed to generate slug:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create post"})
			return
		}

		result, err := db.Exec(
			"INSERT INTO posts (title, content, slug) VALUES (?, ?, ?)",
			payload.Title,
			payload.Content,
			uniqueSlug,
		)
		if err != nil {
			log.Println("failed to create post:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create post"})
			return
		}

		id, _ := result.LastInsertId()
		response := dto.PostResponse{
			ID:      int(id),
			Slug:    uniqueSlug,
			Title:   payload.Title,
			Content: payload.Content,
		}
		c.JSON(http.StatusCreated, response)
	}
}

// GetPostHandler godoc
// @Summary get post by slug
// @Schemes
// @Description fetch a post using its slug
// @Tags posts
// @Accept json
// @Produce json
// @Param slug path string true "Post slug"
// @Success 200 {object} dto.PostResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts/{slug} [get]
func GetPostHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug, ok := parseSlugParam(c)
		if !ok {
			return
		}

		var post dto.PostResponse

		err := db.QueryRow(
			"SELECT COALESCE(id, rowid), slug, title, content, COALESCE(created_at, ''), COALESCE(updated_at, '') FROM posts WHERE slug = ?",
			slug,
		).Scan(&post.ID, &post.Slug, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
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

// GetAllPostsHandler godoc
// @Summary get all posts
// @Schemes
// @Description fetch all blog posts
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {object} []dto.PostResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts [get]
func GetAllPostsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		posts := []dto.PostResponse{}

		rows, err := db.Query(
			"SELECT COALESCE(id, rowid), slug, title, content, COALESCE(created_at, ''), COALESCE(updated_at, '') FROM posts",
		)
		if err != nil {
			log.Println("failed to fetch posts:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch posts"})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var post dto.PostResponse
			if err := rows.Scan(&post.ID, &post.Slug, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt); err != nil {
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

// UpdatePostHandler godoc
// @Summary update post by slug
// @Schemes
// @Description update an existing post using its slug
// @Tags posts
// @Accept json
// @Produce json
// @Param slug path string true "Post slug"
// @Param payload body dto.PostCreateRequest true "Post payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts/{slug} [put]
func UpdatePostHandler(db *sql.DB) gin.HandlerFunc {
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

		var payload dto.PostCreateRequest
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		payload.Title = strings.TrimSpace(payload.Title)
		payload.Content = strings.TrimSpace(payload.Content)

		uniqueSlug, err := utils.GenerateUniquePostSlug(db, payload.Title, existingID)
		if err != nil {
			log.Println("failed to generate slug:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update post"})
			return
		}

		result, err := db.Exec(
			"UPDATE posts SET title = ?, content = ?, slug = ? WHERE id = ?",
			payload.Title,
			payload.Content,
			uniqueSlug,
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

		c.JSON(http.StatusOK, gin.H{"message": "post updated", "slug": uniqueSlug})
	}
}

// DeletePostHandler godoc
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
func DeletePostHandler(db *sql.DB) gin.HandlerFunc {
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

func parseSlugParam(c *gin.Context) (string, bool) {
	slug := strings.TrimSpace(c.Param("slug"))
	if slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid slug"})
		return "", false
	}

	return slug, true
}

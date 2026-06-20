package handlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"simaopsbarbosa/backend/internal/middleware"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	router.POST("/auth/login", LoginHandler())

	// public routes
	router.GET("/posts", GetAllPostsHandler(db))
	router.GET("/posts/:slug", GetPostHandler(db))

	// protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/posts", CreatePostHandler(db))
		protected.PUT("/posts/:slug", UpdatePostHandler(db))
		protected.DELETE("/posts/:slug", DeletePostHandler(db))
		protected.POST("/upload", UploadImageHandler())
	}
}

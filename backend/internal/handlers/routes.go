package handlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	router.POST("/auth/register", RegisterHandler(db))
	router.POST("/auth/login", LoginHandler(db))

	router.POST("/posts", CreatePostHandler(db))
	router.GET("/posts", GetAllPostsHandler(db))
	router.GET("/posts/:slug", GetPostHandler(db))
	router.PUT("/posts/:slug", UpdatePostHandler(db))
	router.DELETE("/posts/:slug", DeletePostHandler(db))
}

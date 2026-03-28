package main

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password,omitempty" binding:"required"`
	CreatedAt string `json:"created_at,omitempty"`
}

type Post struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id" binding:"required"`
	Slug      string `json:"slug,omitempty"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	CreatedAt string `json:"created_at,omitempty"`
}

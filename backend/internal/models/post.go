package models

type Post struct {
	ID        int    `json:"id"`
	Slug      string `json:"slug,omitempty"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

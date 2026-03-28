package models

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password,omitempty" binding:"required"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

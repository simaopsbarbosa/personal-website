package dto

type PostCreateRequest struct {
	UserID  int    `json:"user_id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Slug    string `json:"slug,omitempty"`
}

type PostUpdateRequest struct {
	UserID  int    `json:"user_id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Slug    string `json:"slug,omitempty"`
}

type PostResponse struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

package dto

type PostCreateRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostUpdateRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type PostResponse struct {
	ID        int    `json:"id"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

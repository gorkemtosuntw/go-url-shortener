package model

import "time"

// URL represents the URL entity in the system
type URL struct {
	ID             int64     `json:"id"`
	OriginalURL    string    `json:"original_url"`
	ShortURL       string    `json:"short_url"`
	ShortCode      string    `json:"short_code"`
	CreatedAt      time.Time `json:"created_at"`
	LastAccessedAt time.Time `json:"last_accessed_at,omitempty"`
	ClickCount     int       `json:"click_count"`
}

// CreateURLRequest represents the request for creating a short URL
type CreateURLRequest struct {
	OriginalURL string `json:"original_url"`
}

// CreateURLResponse represents the response after creating a short URL
type CreateURLResponse struct {
	ShortURL string `json:"short_url"`
}

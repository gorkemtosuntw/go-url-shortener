package service

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"url-shortener/internal/model"
	"url-shortener/internal/repository"
)

// URLService defines the interface for URL business operations
type URLService interface {
	CreateShortURL(ctx context.Context, req *model.CreateURLRequest) (*model.CreateURLResponse, error)
	GetOriginalURL(ctx context.Context, shortCode string) (string, error)
}

type urlService struct {
	repo        repository.URLRepository
	baseURL     string
}

// NewURLService creates a new URL service
func NewURLService(repo repository.URLRepository, baseURL string) URLService {
	return &urlService{
		repo:    repo,
		baseURL: baseURL,
	}
}

func (s *urlService) CreateShortURL(ctx context.Context, req *model.CreateURLRequest) (*model.CreateURLResponse, error) {
	// Validate URL
	if _, err := url.ParseRequestURI(req.OriginalURL); err != nil {
		return nil, fmt.Errorf("invalid URL format: %w", err)
	}

	// Generate short code
	shortCode := generateShortCode()
	shortURL := fmt.Sprintf("%s/%s", s.baseURL, shortCode)

	// Create URL entity
	urlEntity := &model.URL{
		OriginalURL: req.OriginalURL,
		ShortURL:    shortURL,
		ShortCode:   shortCode,
		CreatedAt:   time.Now(),
		ClickCount:  0,
	}

	// Save to repository
	if err := s.repo.Create(ctx, urlEntity); err != nil {
		return nil, fmt.Errorf("failed to create URL: %w", err)
	}

	return &model.CreateURLResponse{
		ShortURL: shortURL,
	}, nil
}

func (s *urlService) GetOriginalURL(ctx context.Context, shortCode string) (string, error) {
	// Get URL from repository
	url, err := s.repo.GetByShortCode(ctx, shortCode)
	if err != nil {
		return "", fmt.Errorf("failed to get URL: %w", err)
	}
	if url == nil {
		return "", fmt.Errorf("URL not found")
	}

	// Increment click count
	if err := s.repo.IncrementClickCount(ctx, shortCode); err != nil {
		// Log error but don't fail the request
		fmt.Printf("failed to increment click count: %v\n", err)
	}

	return url.OriginalURL, nil
}

func generateShortCode() string {
	return uuid.New().String()[:8]
}

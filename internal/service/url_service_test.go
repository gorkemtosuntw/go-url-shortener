package service

import (
	"context"
	"testing"
	"time"

	"url-shortener/internal/model"
)

// MockURLRepository is a mock implementation of URLRepository
type MockURLRepository struct {
	createFunc           func(ctx context.Context, url *model.URL) error
	getByShortCodeFunc  func(ctx context.Context, shortCode string) (*model.URL, error)
	incrementClickFunc   func(ctx context.Context, shortCode string) error
}

func (m *MockURLRepository) Create(ctx context.Context, url *model.URL) error {
	return m.createFunc(ctx, url)
}

func (m *MockURLRepository) GetByShortCode(ctx context.Context, shortCode string) (*model.URL, error) {
	return m.getByShortCodeFunc(ctx, shortCode)
}

func (m *MockURLRepository) IncrementClickCount(ctx context.Context, shortCode string) error {
	return m.incrementClickFunc(ctx, shortCode)
}

func TestCreateShortURL(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		mockCreate  func(ctx context.Context, url *model.URL) error
		wantErr     bool
	}{
		{
			name: "valid url",
			url:  "https://example.com",
			mockCreate: func(ctx context.Context, url *model.URL) error {
				return nil
			},
			wantErr: false,
		},
		{
			name: "invalid url",
			url:  "not-a-url",
			mockCreate: func(ctx context.Context, url *model.URL) error {
				return nil
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MockURLRepository{
				createFunc: tt.mockCreate,
			}
			service := NewURLService(repo, "http://localhost:8080")

			_, err := service.CreateShortURL(context.Background(), &model.CreateURLRequest{
				OriginalURL: tt.url,
			})

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateShortURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOriginalURL(t *testing.T) {
	tests := []struct {
		name       string
		shortCode  string
		mockURL    *model.URL
		mockErr    error
		wantURL    string
		wantErr    bool
	}{
		{
			name:      "existing url",
			shortCode: "abc123",
			mockURL: &model.URL{
				OriginalURL: "https://example.com",
				ShortCode:   "abc123",
				CreatedAt:   time.Now(),
			},
			mockErr: nil,
			wantURL: "https://example.com",
			wantErr: false,
		},
		{
			name:      "non-existing url",
			shortCode: "notfound",
			mockURL:   nil,
			mockErr:   nil,
			wantURL:   "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &MockURLRepository{
				getByShortCodeFunc: func(ctx context.Context, shortCode string) (*model.URL, error) {
					return tt.mockURL, tt.mockErr
				},
				incrementClickFunc: func(ctx context.Context, shortCode string) error {
					return nil
				},
			}
			service := NewURLService(repo, "http://localhost:8080")

			got, err := service.GetOriginalURL(context.Background(), tt.shortCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOriginalURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantURL {
				t.Errorf("GetOriginalURL() = %v, want %v", got, tt.wantURL)
			}
		})
	}
}

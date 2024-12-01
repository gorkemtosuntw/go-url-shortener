package repository

import (
	"context"
	"database/sql"
	"time"

	"url-shortener/internal/model"
)

// URLRepository defines the interface for URL storage operations
type URLRepository interface {
	Create(ctx context.Context, url *model.URL) error
	GetByShortCode(ctx context.Context, shortCode string) (*model.URL, error)
	IncrementClickCount(ctx context.Context, shortCode string) error
}

type postgresURLRepository struct {
	db *sql.DB
}

// NewURLRepository creates a new PostgreSQL URL repository
func NewURLRepository(db *sql.DB) URLRepository {
	return &postgresURLRepository{db: db}
}

func (r *postgresURLRepository) Create(ctx context.Context, url *model.URL) error {
	query := `
		INSERT INTO urls (original_url, short_url, short_code, created_at, click_count)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	return r.db.QueryRowContext(
		ctx,
		query,
		url.OriginalURL,
		url.ShortURL,
		url.ShortCode,
		url.CreatedAt,
		url.ClickCount,
	).Scan(&url.ID)
}

func (r *postgresURLRepository) GetByShortCode(ctx context.Context, shortCode string) (*model.URL, error) {
	query := `
		SELECT id, original_url, short_url, short_code, created_at, last_accessed_at, click_count
		FROM urls
		WHERE short_code = $1`

	url := &model.URL{}
	var lastAccessedAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, shortCode).Scan(
		&url.ID,
		&url.OriginalURL,
		&url.ShortURL,
		&url.ShortCode,
		&url.CreatedAt,
		&lastAccessedAt,
		&url.ClickCount,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if lastAccessedAt.Valid {
		url.LastAccessedAt = lastAccessedAt.Time
	}

	return url, nil
}

func (r *postgresURLRepository) IncrementClickCount(ctx context.Context, shortCode string) error {
	query := `
		UPDATE urls
		SET click_count = click_count + 1,
			last_accessed_at = $1
		WHERE short_code = $2`

	_, err := r.db.ExecContext(ctx, query, time.Now(), shortCode)
	return err
}

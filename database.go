package main

import (
	"database/sql"
)

func initDatabase(db *sql.DB) error {
	// Create URLs table if not exists
	query := `
		CREATE TABLE IF NOT EXISTS urls (
			id SERIAL PRIMARY KEY,
			original_url TEXT NOT NULL,
			short_url VARCHAR(255) NOT NULL,
			short_code VARCHAR(50) NOT NULL UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			last_accessed_at TIMESTAMP,
			click_count INTEGER DEFAULT 0
		)`

	_, err := db.Exec(query)
	return err
}

package main

import (
	"database/sql"
	"log"
)

func initDatabase(db *sql.DB) {
	// URLs tablosunu oluştur
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
			id SERIAL PRIMARY KEY,
			original_url TEXT NOT NULL,
			short_url VARCHAR(255) NOT NULL,
			short_code VARCHAR(50) NOT NULL UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			last_accessed_at TIMESTAMP,
			click_count INTEGER DEFAULT 0
		)
	`)

	if err != nil {
		log.Fatal("Veritabanı tablosu oluşturulamadı:", err)
	}
}

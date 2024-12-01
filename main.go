package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"url-shortener/internal/config"
	"url-shortener/internal/handler"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Establish a database connection
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User,
		cfg.Database.Password, cfg.Database.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Initialize database schema
	if err := initDatabase(db); err != nil {
		log.Fatal(err)
	}

	// Initialize application dependencies
	urlRepo := repository.NewURLRepository(db)
	urlService := service.NewURLService(urlRepo, cfg.Server.BaseURL)
	urlHandler := handler.NewURLHandler(urlService)

	// Setup HTTP request router
	r := mux.NewRouter()
	urlHandler.RegisterRoutes(r)

	// Start the server
	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	fmt.Printf("Server is running on %s://%s%s\n",
		cfg.Server.Protocol, cfg.Server.Domain, serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, r))
}

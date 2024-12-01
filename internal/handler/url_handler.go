package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"url-shortener/internal/model"
	"url-shortener/internal/service"
)

// URLHandler handles HTTP requests for URL operations
type URLHandler struct {
	service service.URLService
}

// NewURLHandler creates a new URL handler
func NewURLHandler(service service.URLService) *URLHandler {
	return &URLHandler{
		service: service,
	}
}

// RegisterRoutes registers the URL routes
func (h *URLHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/shorten", h.ShortenURL).Methods(http.MethodPost)
	r.HandleFunc("/{shortCode}", h.RedirectURL).Methods(http.MethodGet)
}

// ShortenURL handles the URL shortening request
func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req model.CreateURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.service.CreateShortURL(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// RedirectURL handles the URL redirection
func (h *URLHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	originalURL, err := h.service.GetOriginalURL(r.Context(), shortCode)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

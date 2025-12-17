package handler

import (
	"cerca-scraper/internal/constants"
	"cerca-scraper/internal/scraper"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	// Add other dependencies here if needed in the future
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleSingleStation(w http.ResponseWriter, r *http.Request) {
	stationName := r.PathValue("stationNameSlug")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	stationSlug, exists := constants.Stations[stationName]
	if !exists {
		http.Error(w, "Station not found", http.StatusNotFound)
		return
	}

	data, err := scraper.ScrapeStation(stationSlug)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error scraping station: %v", err), http.StatusInternalServerError)
		return
	}

	// Removed RabbitMQ publishing - data is now only returned in HTTP response
	// If you need to process this data elsewhere, you can:
	// - Save to database
	// - Call another service directly
	// - Process inline here

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

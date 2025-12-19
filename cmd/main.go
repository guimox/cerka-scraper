package main

import (
	"cerca-scraper/internal/config"
	"cerca-scraper/internal/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /schedule/{stationNameSlug}", handler.HandleSingleStation)
	mux.HandleFunc("GET /schedule/{stationNameSlug}/{trainName}", handler.HandleSingleTrain)
	mux.HandleFunc("GET /stations", handler.HandleAllStations)

	server := &http.Server{
		Addr:         config.GetServerAddress(),
		Handler:      mux,
		ReadTimeout:  config.Server.ReadTimeout,
		WriteTimeout: config.Server.WriteTimeout,
		IdleTimeout:  config.Server.IdleTimeout,
	}

	fmt.Printf("Server starting on http://localhost%s\n", config.GetServerAddress())
	log.Fatal(server.ListenAndServe())
}

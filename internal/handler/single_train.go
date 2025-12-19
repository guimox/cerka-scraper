package handler

import (
	"cerca-scraper/internal/scraper"
	"encoding/json"
	"net/http"
)

type Train struct {
	Time        string `json:"time"`
	Destination string `json:"destination"`
	Name        string `json:"name"`
	Via         string `json:"via"`
}

func findTrainByName(trains []Train, targetName string) *Train {
	for _, t := range trains {
		if t.Name == targetName {
			return &t
		}
	}
	return nil
}

func HandleSingleTrain(w http.ResponseWriter, r *http.Request) {
	trainName := r.PathValue("trainName")
	stationName := r.PathValue("stationNameSlug")

	data, err := scraper.ScrapeStation(stationName)

	if err != nil {
		http.Error(w, "Error al obtener los datos", http.StatusInternalServerError)
		return
	}

	var idx int = -1
	for i, t := range data.Trains {
		if t.TrainName == trainName {
			idx = i
			break
		}
	}

	if idx == -1 {
		http.Error(w, "Tren no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Trains[idx])
}

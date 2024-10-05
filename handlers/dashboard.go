package handlers

import (
	"encoding/json"
	"infotech-calendar/database"
	"net/http"
)

func GetEventData(w http.ResponseWriter, r *http.Request) {
	eventList := database.GetAllEvents()

	jsonData, err := json.Marshal(eventList)
	if err != nil {
		http.Error(w, "Failed to marshal event data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

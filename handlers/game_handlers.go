package handlers

import (
	"encoding/json"
	"hw_server/business"
	"hw_server/model"
	"net/http"
)

// GameHandler validates and parses request, calls business logic.
func GameHandler(w http.ResponseWriter, request *http.Request) {
	if !isValidContentType(request.Header) {
		msg := "Content-Type header is not application/json"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}
	if !isValidSourceType(request.Header) {
		msg := "Source-Type header is not valid"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	sourceType := request.Header.Get("Source-Type")
	event := model.GameEvent{
		SourceType: sourceType,
	}
	err := json.NewDecoder(request.Body).Decode(&event)
	if err != nil {
		msg := "JSON decoding error: " + err.Error()
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	business.HandleGameEvent(event)
}

func isValidSourceType(header http.Header) bool {
	sourceType := header.Get("Source-Type")
	if sourceType == "game" || sourceType == "server" || sourceType == "payment" {
		return true
	}
	return false
}

func isValidContentType(header http.Header) bool {
	if header.Get("Content-Type") == "application/json" {
		return true
	}

	return false
}

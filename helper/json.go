package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode == 500 {
		log.Println("Responding with 500 error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, statusCode, errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

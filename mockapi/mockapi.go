package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Data string `json:"data"`
}

func main() {
	http.HandleFunc("/api", handleRequest)
	log.Println("Mock API server started on: http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		Message string `json:"message"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		return
	}

	response := Response{}

	// Predetermined responses
	switch requestData.Message {
	case "Hello":
		response.Data = "Hi there!"
	case "How are you?":
		response.Data = "I'm just a mock server, but I'm doing great!"
	case "Morgen":
		response.Data = "Guten Morgen!"
	default:
		response.Data = "I don't understand that question."
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

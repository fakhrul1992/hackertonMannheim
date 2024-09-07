package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type APIData struct {
	Data string `json:"data"`
}

func main() {
	http.HandleFunc("/api/send", handleAPIRequest)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/* old
func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
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

	apiResponse, err := sendDataToAPI(requestData.Message)
	if err != nil {
		http.Error(w, "Error calling API", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(apiResponse)
}
*/

func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Message string `json:"message"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		return
	}

	apiResponse, err := sendDataToAPI(requestData.Message)
	if err != nil {
		log.Printf("Error sending data to API: %v", err)
		http.Error(w, "Error communicating with the API", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(apiResponse)
	// Process and send the API response back to the client...
}

/////////////////////////////////

/*  old
func sendDataToAPI(message string) (APIData, error) {
	apiURL := "https://example.com/api" // Replace with your REST API URL
	formData := url.Values{"message": {message}}

	resp, err := http.PostForm(apiURL, formData)
	if err != nil {
		return APIData{}, err
	}
	defer resp.Body.Close()

	var apiData APIData
	err = json.NewDecoder(resp.Body).Decode(&apiData)
	if err != nil {
		return APIData{}, err
	}

	return apiData, nil
}
*/

/*  old 2
func sendDataToAPI(message string) (APIData, error) {
	//apiURL := "http://localhost:8081/api" // URL for the mock API
	apiURL := "http://mockapi:8081/api"
	formData := url.Values{"message": {message}}

	resp, err := http.PostForm(apiURL, formData)
	if err != nil {
		return APIData{}, err
	}
	defer resp.Body.Close()

	var apiData APIData
	err = json.NewDecoder(resp.Body).Decode(&apiData)
	if err != nil {
		return APIData{}, err
	}

	return apiData, nil
}
*/

func sendDataToAPI(message string) (APIData, error) {
	apiURL := "http://mockapi:8081/api" // Ensure this is correct
	jsonData := map[string]string{"message": message}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return APIData{}, err
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return APIData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return APIData{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var apiData APIData
	err = json.NewDecoder(resp.Body).Decode(&apiData)
	if err != nil {
		return APIData{}, err
	}

	return apiData, nil
}

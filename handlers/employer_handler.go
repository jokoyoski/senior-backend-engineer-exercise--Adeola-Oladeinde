// handlers/employee_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"syndio-backend-app/models"
)

func SaveJobsHandler(w http.ResponseWriter, r *http.Request) {
	var jobs []models.Job
	if err := json.NewDecoder(r.Body).Decode(&jobs); err != nil {
		http.Error(w, "Invalid job data", http.StatusBadRequest)
		return
	}

	for _, job := range jobs {
		if err := job.Save(); err != nil {
			http.Error(w, "Error saving job data", http.StatusInternalServerError)
			return
		}
	}

	// Create a response object
	response := map[string]string{
		"message": "Job data saved successfully",
	}

	// Set the response header to indicate it's JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Write the JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

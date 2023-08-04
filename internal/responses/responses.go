package responses

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

func ErrorResponses(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := ErrorResponse{
		StatusCode: status,
		Message:    http.StatusText(status),
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  interface{} `json:"result,omitempty"`
}

func WriteResponse(w http.ResponseWriter, statusCode int, message string, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(APIResponse{
		Status:  statusCode,
		Message: message,
		Result: result,
	})
}

package response

import (
	"encoding/json"
	"net/http"
)

func Encode(w http.ResponseWriter, body any, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

package response

import (
	"encoding/json"
	"net/http"
)

func Encode(w http.ResponseWriter, body any, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := Response{
		Data:       body,
		Error:      "",
		StatusCode: statusCode,
	}

	if err, implement := body.(error); implement {
		resp.Error = err.Error()
		resp.Data = nil
	}

	json.NewEncoder(w).Encode(resp)
}

type Response struct {
	Data       any    `json:"data"`
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

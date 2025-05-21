package backend

import (
	"encoding/json"
	"net/http"
)

const failedToEncodeResponse = `{"status": "FAILURE", "message": "Failed to encode response.", "data": {}}`

func write(w http.ResponseWriter, status int, msg interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(msg); err != nil {
		http.Error(w, failedToEncodeResponse, http.StatusInternalServerError)
		return
	}
}

func must_read_body(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	decoder := json.NewDecoder(r.Body)
	// decoder.DisallowUnknownFields() // Prevents unknown fields in the request
	if err := decoder.Decode(v); err != nil {
		http.Error(w, `{"status": "FAILURE", "message": "Invalid request format.", "data": {}}`, http.StatusBadRequest)
		return false
	}
	return true
}

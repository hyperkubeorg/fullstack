package backend

import (
	"encoding/json"
	"net/http"
	"time"
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

func writeLoginRequired(w http.ResponseWriter) {
	write(w, http.StatusUnauthorized,
		struct {
			Status  string   `json:"status"`
			Message string   `json:"message"`
			Data    struct{} `json:"data"`
		}{
			Status:  "FAILURE",
			Message: "User is not logged in.",
			Data:    struct{}{},
		},
	)
}

func writeCookie(w http.ResponseWriter, name string, value string, ttl_in_seconds int) {
	http.SetCookie(w, &http.Cookie{
		Name:    name,
		Value:   value,
		Path:    "/",
		Expires: time.Now().Add(time.Duration(ttl_in_seconds) * time.Second),
	})
}

func deleteCookie(w http.ResponseWriter, name string) {
	http.SetCookie(w, &http.Cookie{
		Name:    name,
		Value:   "",
		Path:    "/",
		Expires: time.Now().Add(-time.Hour),
	})
}

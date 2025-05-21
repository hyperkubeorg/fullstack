package backend

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func AddRoutes(r *mux.Router) (*mux.Router, error) {
	r.HandleFunc("/api/v1/time", timeHandler).Methods("GET")

	return r, nil
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	// serialize and send the response
	write(w, http.StatusOK,
		struct {
			Status  string      `json:"status"`
			Message string      `json:"message"`
			Data    interface{} `json:"data"`
		}{
			Status:  "SUCCESS",
			Message: "User created successfully.",
			Data: map[string]interface{}{
				"current_time_utc": time.Now().UTC().Format(time.RFC3339),
			},
		},
	)
}

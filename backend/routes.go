package backend

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func AddRoutes(r *mux.Router) (*mux.Router, error) {
	r.HandleFunc("/api/v1/time", timeHandler).Methods("GET")
	r.HandleFunc("/api/v1/auth/signup", signupHandler).Methods("POST")
	r.HandleFunc("/api/v1/auth/login", loginHandler).Methods("POST")

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
			Message: "Time has been retrieved.",
			Data: map[string]interface{}{
				"current_time_utc": time.Now().UTC().Format(time.RFC3339),
			},
		},
	)
}

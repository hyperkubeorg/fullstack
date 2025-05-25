package backend

import (
	"net/http"

	"github.com/hyperkubeorg/fullstack/models"
)

func authLogoutHandler(w http.ResponseWriter, r *http.Request) {
	models.DestroyUserSessionFromRequest(r)

	deleteCookie(w, "session_token")

	// serialize and send the response
	write(w, http.StatusOK,
		struct {
			Status  string   `json:"status"`
			Message string   `json:"message"`
			Data    struct{} `json:"data"`
		}{
			Status:  "SUCCESS",
			Message: "Logged out successfully.",
			Data:    struct{}{},
		},
	)
}

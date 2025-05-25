package backend

import (
	"net/http"

	"github.com/hyperkubeorg/fullstack/models"
)

func authStatusHandler(w http.ResponseWriter, r *http.Request) {
	user, err := models.GetUserFromRequest(r)
	if err != nil || user == nil {
		writeLoginRequired(w)
		return
	}

	// serialize and send the response
	write(w, http.StatusOK,
		struct {
			Status  string      `json:"status"`
			Message string      `json:"message"`
			Data    interface{} `json:"data"`
		}{
			Status:  "SUCCESS",
			Message: "User is logged in as " + user.Name,
			Data: struct {
				Username string `json:"username"`
				IsAdmin  bool   `json:"is_admin"`
				IsBanned bool   `json:"is_banned"`
			}{
				Username: user.Name,
				IsAdmin:  user.IsAdmin,
				IsBanned: user.IsBanned,
			},
		},
	)
}

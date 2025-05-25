package backend

import (
	"net/http"

	"github.com/hyperkubeorg/fullstack/models"
)

func authSignupHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Username        string `json:"username"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		TermsAccepted   bool   `json:"terms"`
		PrivacyAccepted bool   `json:"privacy"`
	}
	if !must_read_body(w, r, &request) {
		return
	}

	if !request.TermsAccepted {
		write(w, http.StatusBadRequest,
			struct {
				Status  string   `json:"status"`
				Message string   `json:"message"`
				Data    struct{} `json:"data"`
			}{
				Status:  "FAILURE",
				Message: "You must accept the Terms of Service.",
				Data:    struct{}{},
			},
		)
		return
	}

	if !request.PrivacyAccepted {
		write(w, http.StatusBadRequest,
			struct {
				Status  string   `json:"status"`
				Message string   `json:"message"`
				Data    struct{} `json:"data"`
			}{
				Status:  "FAILURE",
				Message: "You must accept the Privacy Policy.",
				Data:    struct{}{},
			},
		)
		return
	}

	new_user := &models.User{
		Name:            request.Username,
		Email:           request.Email,
		Password:        request.Password,
		PasswordConfirm: request.Password,
	}

	db := models.GetDB()
	txn := db.Create(new_user)
	if txn.Error != nil {
		write(w, http.StatusInternalServerError,
			struct {
				Status  string   `json:"status"`
				Message string   `json:"message"`
				Data    struct{} `json:"data"`
			}{
				Status:  "FAILURE",
				Message: "Failed to create user: " + txn.Error.Error(),
				Data:    struct{}{},
			},
		)
		return
	}

	sessionToken := &models.UserSession{
		UserID: new_user.ID,
	}

	if err := db.Create(sessionToken).Error; err != nil {
		write(w, http.StatusInternalServerError,
			struct {
				Status  string   `json:"status"`
				Message string   `json:"message"`
				Data    struct{} `json:"data"`
			}{
				Status:  "FAILURE",
				Message: "Failed to create session token: " + err.Error(),
				Data:    struct{}{},
			},
		)
		return
	}

	writeCookie(w, "session_token", sessionToken.Token, 2592000)

	// serialize and send the response
	write(w, http.StatusOK,
		struct {
			Status  string   `json:"status"`
			Message string   `json:"message"`
			Data    struct{} `json:"data"`
		}{
			Status:  "SUCCESS",
			Message: "User created successfully.",
			Data:    struct{}{},
		},
	)
}

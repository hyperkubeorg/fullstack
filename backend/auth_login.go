package backend

import (
	"crypto/sha512"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/hyperkubeorg/fullstack/models"
)

func authLoginHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		UsernameOrEmail string `json:"username_or_email"`
		Password        string `json:"password"`
	}
	if !must_read_body(w, r, &request) {
		return
	}

	user := &models.User{}

	if _, err := mail.ParseAddress(request.UsernameOrEmail); err == nil {
		hasher := sha512.New()
		hasher.Write([]byte(request.UsernameOrEmail))
		emailHash := fmt.Sprintf("%x", hasher.Sum(nil))
		models.GetDB().Where("email_hash = ?", emailHash).First(user)
	} else {
		models.GetDB().Where("name = ?", request.UsernameOrEmail).First(user)
	}

	if user.ID == "" || !user.IsValidPassword(request.Password) {
		write(w, http.StatusUnauthorized,
			struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{
				Status:  "ERROR",
				Message: "Login failed.",
			},
		)
		return
	}

	db := models.GetDB()

	sessionToken := &models.UserSession{
		UserID: user.ID,
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
			Message: "Login Successful.",
			Data:    struct{}{},
		},
	)
}

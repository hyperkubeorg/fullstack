package models

import (
	"crypto/rand"
	"fmt"
	"net/http"
)

// Helper function to generate a random hex string of the specified length
func randomHex(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", bytes), nil
}

func GetUserFromRequest(r *http.Request) (*User, error) {
	// get the session cookie
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// no cookie found, user is not logged in
			return nil, nil
		}
		// some other error occurred
		return nil, err
	}

	sessionToken := cookie.Value
	db := GetDB()

	// Fetch the UserSession using the session token
	var session UserSession
	if err := db.Where("token = ?", sessionToken).First(&session).Error; err != nil {
		return nil, nil
	}

	// Fetch the associated User using the UserID from the session
	var user User
	if err := db.Where("id = ?", session.UserID).First(&user).Error; err != nil {
		return nil, nil
	}

	if user.IsBanned {
		return nil, fmt.Errorf("user is banned")
	}

	return &user, nil
}

func DestroyUserSessionFromRequest(r *http.Request) error {
	// get the session cookie
	cookie, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// no cookie found, user is not logged in
			return nil
		}
		// some other error occurred
		return err
	}

	sessionToken := cookie.Value
	db := GetDB()

	// Delete the UserSession using the session token
	if err := db.Where("token = ?", sessionToken).Delete(&UserSession{}).Error; err != nil {
		return err
	}

	return nil
}

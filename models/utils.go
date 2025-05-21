package models

import (
	"crypto/rand"
	"fmt"
)

// Helper function to generate a random hex string of the specified length
func randomHex(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", bytes), nil
}

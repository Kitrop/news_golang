package utils

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/matthewhartstonge/argon2"
)

// HashPassword hashes the password using Argon2
func HashPassword(password string) (string, error) {
	hash, err := argon2Instance.HashEncoded([]byte(password))
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hash), nil
}

// CheckPasswordHash verifies the password against the Argon2 hash
func CheckPasswordHash(password, hash string) bool {
	ok, err := argon2.VerifyEncoded([]byte(password), []byte(hash))

	if err != nil {
		return false
	}

	return ok
}

// Improved password strength check with comments
func IsStrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasUpper, hasLower, hasDigit, hasSpecial := false, false, false, false
	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(string(ch)):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasDigit && hasSpecial
}

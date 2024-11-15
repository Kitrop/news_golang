package utils

import (
	"errors"
	"fmt"
	"news-go/config"
	"news-go/models"
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckEmailExist(email string) bool {
	var findByEmail = models.User{Email: email}
	result := config.DB.First(&findByEmail)

	if err := result.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}

	return false
}

func CheckUsernameExist(username string) bool {
	var findByUsername = models.User{Username: username}
	result := config.DB.First(&findByUsername)

	if err := result.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}

	return false
}

func IsStrongPassword(password string) bool {
	const minLength = 8

	if len(password) < minLength {
		return false
	}

	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)

	return re.MatchString(email)
}

func ValidateUserInput(input models.User) error {
	if CheckUsernameExist(input.Username) {
		return fmt.Errorf("this username already exists")
	}
	if !IsValidEmail(input.Email) {
		return fmt.Errorf("email is invalid")
	}
	if CheckEmailExist(input.Email) {
		return fmt.Errorf("this email already exists")
	}
	if !IsStrongPassword(input.Password) {
		return fmt.Errorf("this password is not strong")
	}
	return nil
}

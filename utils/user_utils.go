package utils

import (
	"fmt"
	"news-go/config"
	"news-go/models"
	"regexp"

	"github.com/matthewhartstonge/argon2"
)

var argon2Instance = argon2.DefaultConfig()

// Checks if a user with the given email already exists in the database
func CheckEmailExist(email string) (bool, error) {
	var count int64
	err := config.DB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}


// Checks if a user with the given username already exists in the database
func CheckUsernameExist(username string) (bool, error) {
	var count int64
	err := config.DB.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Improved email validation with a compiled regex outside function
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// Сhecks if the given email address is valid based on a regular expression
func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// Validates user input for creating a new user
// It checks if the username and email already exist and if the password meets certain criteria
func ValidateUserInput(input models.User) error {
	usernameExists, err := CheckUsernameExist(input.Username)

	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("some error username")
	}

	if usernameExists {
		return fmt.Errorf("username %s already exists", input.Username)
	}

	if !IsValidEmail(input.Email) {
		return fmt.Errorf("invalid email address")
	}

	emailExists, err := CheckEmailExist(input.Email)

	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("some error email")
	}

	if emailExists {
		return fmt.Errorf("email %s already exists", input.Email)
	}

	if !IsStrongPassword(input.Password) {
		return fmt.Errorf("password is not strong enough")
	}
	return nil
}

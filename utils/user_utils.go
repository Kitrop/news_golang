package utils

import (
	"fmt"
	"news-go/config"
	"news-go/models"
	"regexp"

	"github.com/matthewhartstonge/argon2"
)

var argon2Instance = argon2.DefaultConfig()

func CheckEmailExist(email string) (bool, error) {
	var count int64
	err := config.DB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

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

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

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

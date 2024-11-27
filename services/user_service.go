package services

import (
	"errors"
	"news-go/models"
	"news-go/repositories"
	"news-go/utils"
)

// CreateUser creates a new user
func CreateUser(input *models.User) (map[string]string, string, error) {
	// Валидация входных данных
	if err := utils.ValidateUserInput(*input); err != nil {
		return nil, "", err
	}

	// Hash the password
	hashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, "", errors.New("invalid password, try another input password")
	}

	input.Password = hashPassword

	// Create a new user
	newUserData, err := repositories.CreateUserInDB(input)
	if err != nil {
		return nil, "", errors.New("failed to create user in database")
	}

	accessToken, err := utils.GenerateJWT(newUserData.ID, newUserData.Username, newUserData.Email, string(newUserData.Role))

	if err != nil {
		return nil, "", errors.New(err.Error())
	}

	// Return user data
	userData := map[string]string{
		"username": input.Username,
		"email":    input.Email,
		"role":     string(input.Role),
	}
	return userData, accessToken, nil
}

// GetAllUsers returns all users
func GetAllUsers() ([]models.User, error) {
	return repositories.FindAllUsers()
}

// LoginUser performs user authentication and generates a JWT
func LoginUser(username, password string) (map[string]string, string, error) {
	// Find the user by username
	user, err := repositories.FindUserByUsername(username)
	if err != nil {
		return nil, "", errors.New("user not found")
	}

	// Check the password
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, "", errors.New("invalid password")
	}

	// Generate JWT
	accessToken, err := utils.GenerateJWT(user.ID, user.Username, user.Email, string(user.Role))
	if err != nil {
		return nil, "", errors.New(err.Error())
	}

	userData := map[string]string{
		"userID": string(rune(user.ID)),
		"username": user.Username,
		"email":    user.Email,
		"role":     string(user.Role),
	}
	return userData, accessToken, nil
}

// ChangeUserPassword changes the user's password
func ChangeUserPassword(accessToken, oldPassword, newPassword string) error {
	// Validate the JWT
	claims, err := utils.ValidateJWT(accessToken)
	if err != nil {
		return errors.New("invalid access token")
	}

	// Find the user by email
	user, err := repositories.FindUserByEmail(claims.Email)
	if err != nil {
		return errors.New("user not found")
	}

	// Check the old password
	if !utils.CheckPasswordHash(oldPassword, user.Password) {
		return errors.New("incorrect old password")
	}

	// Check the strength of the new password
	if !utils.IsStrongPassword(newPassword) {
		return errors.New("new password is weak")
	}

	// Hash the new password
	newHashPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New("failed to hash new password")
	}

	// Update the password in the database
	data := repositories.UpdateUserPassword(user.ID, newHashPassword)

	return data
}

// CreateNewSession creates information about current sessions in the DB
func CreateNewSession(userID uint, ip string, browser string) error {
	if err := repositories.AddActiveSession(userID, ip, browser); err != nil {
		return errors.New(err.Error())
	}

	return nil
}
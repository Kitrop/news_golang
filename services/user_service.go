package services

import (
	"errors"
	"news-go/models"
	"news-go/repositories"
	"news-go/utils"
)

// CreateUser создает нового пользователя
func CreateUser(input *models.User) (map[string]string, string, error) {
	// Валидация входных данных
	if err := utils.ValidateUserInput(*input); err != nil {
		return nil, "", err
	}

	// Хэширование пароля
	hashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, "", errors.New("invalid password, try another input password")
	}

	// Создание нового пользователя
	input.Password = hashPassword
	if err := repositories.CreateUserInDB(input); err != nil {
		return nil, "", errors.New("failed to create user in database")
	}

	// Генерация JWT
	accessToken, err := utils.GenerateJWT(input.Username, input.Email, string(input.Role))
	if err != nil {
		return nil, "", errors.New("failed to generate JWT")
	}

	// Возвращаем данные пользователя
	userData := map[string]string{
		"username": input.Username,
		"email":    input.Email,
		"role":     string(input.Role),
	}
	return userData, accessToken, nil
}

// GetAllUsers возвращает всех пользователей
func GetAllUsers() ([]models.User, error) {
	return repositories.FindAllUsers()
}

// LoginUser выполняет аутентификацию пользователя и генерирует JWT
func LoginUser(username, password string) (map[string]string, string, error) {
	// Поиск пользователя по имени
	user, err := repositories.FindUserByUsername(username)
	if err != nil {
		return nil, "", errors.New("user not found")
	}

	// Проверка пароля
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, "", errors.New("invalid password")
	}

	// Генерация JWT
	accessToken, err := utils.GenerateJWT(user.Username, user.Email, string(user.Role))
	if err != nil {
		return nil, "", errors.New("failed to generate JWT")
	}

	userData := map[string]string{
		"username": user.Username,
		"email":    user.Email,
		"role":     string(user.Role),
	}
	return userData, accessToken, nil
}

// ChangeUserPassword изменяет пароль пользователя
func ChangeUserPassword(accessToken, oldPassword, newPassword string) error {
	// Проверка токена JWT
	claims, err := utils.ValidateJWT(accessToken)
	if err != nil {
		return errors.New("invalid access token")
	}

	// Поиск пользователя по email
	user, err := repositories.FindUserByEmail(claims.Email)
	if err != nil {
		return errors.New("user not found")
	}

	// Проверка старого пароля
	if !utils.CheckPasswordHash(oldPassword, user.Password) {
		return errors.New("incorrect old password")
	}

	// Проверка силы нового пароля
	if !utils.IsStrongPassword(newPassword) {
		return errors.New("new password is weak")
	}

	// Хэширование нового пароля
	newHashPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New("failed to hash new password")
	}

	// Обновление пароля в базе данных
	return repositories.UpdateUserPassword(user.ID, newHashPassword)
}

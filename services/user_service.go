package services

import (
	"errors"
	"fmt"
	"news-go/models"
	"news-go/repositories"
	"news-go/utils"
	"time"
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

	input.Password = hashPassword

	// Создание нового пользователя
	newUserData, err := repositories.CreateUserInDB(input)
	if err != nil {
		return nil, "", errors.New("failed to create user in database")
	}

	accessToken, err := utils.GenerateJWT(newUserData.ID, newUserData.Username, newUserData.Email, string(newUserData.Role))

	if err != nil {
		return nil, "", errors.New(err.Error())
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
	accessToken, err := utils.GenerateJWT(user.ID, user.Username, user.Email, string(user.Role))
	if err != nil {
		return nil, "", errors.New(err.Error())
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
	startTime := time.Now()
	claims, err := utils.ValidateJWT(accessToken)
	if err != nil {
		return errors.New("invalid access token")
	}
	duration := time.Since(startTime).Seconds() * 1000 // миллисекунды
	fmt.Println(duration)

	// Поиск пользователя по email
	startTime1 := time.Now()
	user, err := repositories.FindUserByEmail(claims.Email)
	if err != nil {
		return errors.New("user not found")
	}
	duration1 := time.Since(startTime1).Seconds() * 1000 // миллисекунды
	fmt.Println(duration1)

	// Проверка старого пароля
	startTime2 := time.Now()
	if !utils.CheckPasswordHash(oldPassword, user.Password) {
		return errors.New("incorrect old password")
	}
	duration2 := time.Since(startTime2).Seconds() * 1000 // миллисекунды
	fmt.Println(duration2)

	// Проверка силы нового пароля
	startTime3 := time.Now()
	if !utils.IsStrongPassword(newPassword) {
		return errors.New("new password is weak")
	}
	duration3 := time.Since(startTime3).Seconds() * 1000 // миллисекунды
	fmt.Println(duration3)

	// Хэширование нового пароля
	startTime4 := time.Now()
	newHashPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New("failed to hash new password")
	}
	duration4 := time.Since(startTime4).Seconds() * 1000 // миллисекунды
	fmt.Println(duration4)

	// Обновление пароля в базе данных
	startTime5 := time.Now()
	data := repositories.UpdateUserPassword(user.ID, newHashPassword)
	duration5 := time.Since(startTime5).Seconds() * 1000 // миллисекунды
	fmt.Println(duration5)

	return data
}

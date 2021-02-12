package utils

import (
	"os"
)

// ValidLogin проверка валидности логина
func ValidLogin(login string) bool {
	// re := regexp.MustCompile(`[\w_]{5,}`)
	return false
}

// ValidName проверка валидности имени
func ValidName(name string) bool {
	return false
}

// ValidPassword проверка валидности пароля
func ValidPassword(password string) bool {
	return false
}

// ValidateUserData проверка валидности всех параметров пользователя
func ValidateUserData(userData map[string]string) bool {
	login, ok := userData["login"]
	if !ok {
		return false
	}
	name, ok := userData["name"]
	if !ok {
		return false
	}
	passw, ok := userData["password"]
	if !ok {
		return false
	}

	return ValidLogin(login) && ValidName(name) && ValidPassword(passw)
}

// FileExist проверка существования файла
func FileExist(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

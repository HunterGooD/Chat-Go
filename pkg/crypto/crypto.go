package crypto

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

// Base64Decode из base64 декодирует в строку
func Base64Decode(data string) string {
	result, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return ""
	}
	return string(result)
}

// Base64Encode строку в base64 кодирует
func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// HashPassword создает хэш пароля bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash верификация пароля с хэшом
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

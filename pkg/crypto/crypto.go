package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"math"

	"golang.org/x/crypto/bcrypt"
)

// GenerateKey ключ из до нужной длины
func GenerateKey(key string, size int) []byte {
	var result string
	if len(key) >= size {
		return []byte(key[:size])
	}
	for i := 0; i < int(math.Ceil(float64(size)/float64(len(key)))); i++ {
		result += key
	}
	return []byte(result[:size])
}

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

// EnctyprAES ..
func EnctyprAES(key, data []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	data = paddingPKCS5(data, blockSize)
	cipherText := make([]byte, blockSize+len(data))
	iv := cipherText[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], data)
	return cipherText
}

// DecryptAES ...
func DecryptAES(key, data []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	blockSize := block.BlockSize()
	if len(data) < blockSize {
		return nil
	}
	iv := data[:blockSize]
	data = data[blockSize:]
	if len(data)%blockSize != 0 {
		return nil
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)
	return unpaddingPKS5(data)
}

func unpaddingPKS5(data []byte) []byte {
	length := len(data)
	if length == 0 {
		return nil
	}
	unpadding := int(data[length-1])
	if length < unpadding {
		return nil
	}
	return data[:(length - unpadding)]
}

func paddingPKCS5(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	paddText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, paddText...)
}

package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

var (
	nameLogin = regexp.MustCompile(`^([A-z]+([-_]?[A-z0-9]+){0,2}){4,32}$`)
)

// ValidLogin проверка валидности логина
func ValidLogin(login string) bool {
	return nameLogin.Match([]byte(login))
}

// ValidName проверка валидности имени
func ValidName(name string) bool {
	return nameLogin.Match([]byte(name))
}

// ValidPassword проверка валидности пароля
func ValidPassword(password string) bool {
	return CheckPasswordLever(password) == nil
}

//CheckPasswordLever Сложность пароля
func CheckPasswordLever(ps string) error {
	if len(ps) < 9 {
		return fmt.Errorf("password len is < 9")
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return fmt.Errorf("password need num :%v", err)
	}
	if b, err := regexp.MatchString(a_z, ps); !b || err != nil {
		return fmt.Errorf("password need a_z :%v", err)
	}
	if b, err := regexp.MatchString(A_Z, ps); !b || err != nil {
		return fmt.Errorf("password need A_Z :%v", err)
	}
	if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		return fmt.Errorf("password need symbol :%v", err)
	}
	return nil
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

// ToJSON в JSON
func ToJSON(data interface{}) ([]byte, error) {
	JSONdata, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return JSONdata, nil
}

// FromJSON из JSON
func FromJSON(data []byte, i interface{}) (interface{}, error) {
	if err := json.Unmarshal(data, &i); err != nil {
		return nil, err
	}
	return i, nil
}

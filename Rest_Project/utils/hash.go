package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	cost_value := 14
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost_value)
	return string(bytes), err
}

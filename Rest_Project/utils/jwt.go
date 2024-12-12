package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// This should go in the manifest for the running API, keep that in mind for later
const secretKey = "secretKeyDummyValue"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

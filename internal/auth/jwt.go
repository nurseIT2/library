package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// JwtSecret секретный ключ для JWT токенов
var JwtSecret = "your-secret-key"

// Приватная переменная для использования внутри пакета
var secret = []byte(JwtSecret)

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

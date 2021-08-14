package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// Набор полей для токена
type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func NewClaims(expireDuration time.Duration, username string) *Claims {
	return &Claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		username,
	}
}

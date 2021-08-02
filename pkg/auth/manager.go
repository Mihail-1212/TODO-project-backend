package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mihail-1212/todo-project-backend/pkg/auth"
)

// Структура для авторизации пользователей
type Authorizer struct {
	repo auth.Repository

	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

// Набор полей для токена
type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func SignIn() {

}

func test() {

}

// Авторизация GoLang

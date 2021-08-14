package auth

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mihail-1212/todo-project-backend/pkg/auth/models"
	repo "github.com/mihail-1212/todo-project-backend/pkg/auth/repository"
)

// Структура для авторизации пользователей
type Authorizer struct {
	repo *repo.Repository

	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthorizer(repo *repo.Repository, hashSalt string, signingKey []byte, expireDuration time.Duration) *Authorizer {
	return &Authorizer{
		repo:           repo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: expireDuration,
	}
}

// Генерация пароля из строки пароля и соли
func (a *Authorizer) GeneratePasswordHash(password string) string {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	return fmt.Sprintf("%x", pwd.Sum(nil))
}

func (a *Authorizer) SignInReturnToken(user *models.User) (string, error) {
	user.Password = a.GeneratePasswordHash(user.Password)
	user, err := a.repo.Auth.Get(user.Username, user.Password, a.GeneratePasswordHash)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, NewClaims(a.expireDuration, user.Username))
	return token.SignedString(a.signingKey)
}

func (a *Authorizer) SignUpReturnToken(user *models.User) (string, error) {
	user.Password = a.GeneratePasswordHash(user.Password)

	err := a.repo.Auth.Insert(user, a.GeneratePasswordHash)

	if err != nil {
		return "", err
	}

	return a.SignInReturnToken(user)
}

func (a *Authorizer) ParseTokenReturnUsername(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, err := t.Method.(*jwt.SigningMethodHMAC); !err {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, err := token.Claims.(*Claims); err && token.Valid {
		return claims.Username, nil
	}

	return "", ErrInvalidAccessToken
}

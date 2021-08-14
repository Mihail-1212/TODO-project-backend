package repository

import (
	"github.com/mihail-1212/todo-project-backend/pkg/auth/models"
)

type Authorizer interface {
	Insert(user *models.User, generatePasswordHash func(string) string) error
	Get(username, password string, generatePasswordHash func(string) string) (*models.User, error)
}

type Repository struct {
	Auth Authorizer
}

func NewRepository(auth Authorizer) *Repository {
	return &Repository{
		Auth: auth,
	}
}

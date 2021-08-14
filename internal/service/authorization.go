package service

import (
	repo "github.com/mihail-1212/todo-project-backend/internal/repository"
	"github.com/mihail-1212/todo-project-backend/pkg/domain"
)

type Authorization struct {
	repo *repo.Repository
}

func NewAuthorization(repo *repo.Repository) *Authorization {
	return &Authorization{
		repo: repo,
	}
}

func (a *Authorization) CreateUser(user domain.User) (int, error) {
	return a.repo.User.CreateUser(user)
}

package service

import (
	repo "github.com/mihail-1212/todo-project-backend/internal/repository"
	"github.com/mihail-1212/todo-project-backend/pkg/domain"
)

type User struct {
	repo *repo.Repository
}

func NewUser(repo *repo.Repository) *User {
	return &User{
		repo: repo,
	}
}

func (u *User) IsUserExists(id int) (bool, error) {
	return u.repo.User.IsUserExists(id)
}

func (u *User) DeleteUser(id int) error {
	return u.repo.User.DeleteUser(id)
}

func (u *User) GetUserByID(id int) (*domain.User, error) {
	return u.repo.User.GetUserByID(id)
}

func (u *User) GetUserByLogin(login string) (*domain.User, error) {
	return u.repo.User.GetUserByLogin(login)
}

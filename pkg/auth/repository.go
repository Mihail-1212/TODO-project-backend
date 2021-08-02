package auth

import (
	"context"

	"github.com/mihail-1212/todo-project-backend/pkg/domain"
)

type Repository interface {
	Insert(ctx context.Context, user *domain.User) error
	Get(ctx context.Context, username, password string) (*domain.User, error)
}

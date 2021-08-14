package service

import (
	repo "github.com/mihail-1212/todo-project-backend/internal/repository"
	"github.com/mihail-1212/todo-project-backend/pkg/domain"
)

type AuthorizationService interface {
	CreateUser(domain.User) (int, error)
}

type TaskService interface {
	CreateTask(*domain.Task) (*domain.Task, error)
	UpdateTaskOfUser(*domain.Task) error
	DeleteUserTask(int, *domain.User) error
	GetTaskListOfUser(userId int) ([]domain.Task, error)
	GetUserTaskByID(int, *domain.User) (*domain.Task, error)
}

type UserService interface {
	IsUserExists(int) (bool, error)
	DeleteUser(int) error
	GetUserByID(int) (*domain.User, error)
	GetUserByLogin(string) (*domain.User, error)
}

type Services struct {
	TaskService
	UserService
	AuthorizationService
}

func NewServices(repo *repo.Repository) *Services {
	return &Services{
		TaskService:          NewTask(repo),
		UserService:          NewUser(repo),
		AuthorizationService: NewAuthorization(repo),
	}
}

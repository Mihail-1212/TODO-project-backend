package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mihail-1212/todo-project-backend/internal/repository/postgres"
	"github.com/mihail-1212/todo-project-backend/pkg/auth/models"
	"github.com/mihail-1212/todo-project-backend/pkg/domain"
)

type User interface {
	CreateUser(domain.User) (int, error)
	IsUserExists(int) (bool, error)
	DeleteUser(int) error
	Insert(user *models.User, generatePasswordHash func(string) string) error
	Get(username, password string, generatePasswordHash func(string) string) (*models.User, error)
	GetUserByID(int) (*domain.User, error)
	GetUserByLogin(string) (*domain.User, error)
}

type Task interface {
	CreateTask(*domain.Task) (*domain.Task, error)
	UpdateTaskOfUser(*domain.Task) error
	DeleteUserTask(int, *domain.User) error
	GetTaskListOfUser(userId int) ([]domain.Task, error)
	GetUserTaskByID(int, *domain.User) (*domain.Task, error)
}

type Repository struct {
	User
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Task: postgres.NewTaskPostgres(db),
		User: postgres.NewUserPostgres(db),
	}
}

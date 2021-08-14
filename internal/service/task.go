package service

import (
	repo "github.com/mihail-1212/todo-project-backend/internal/repository"
	"github.com/mihail-1212/todo-project-backend/pkg/domain"
)

type Task struct {
	repo *repo.Repository
}

func NewTask(repo *repo.Repository) *Task {
	return &Task{
		repo: repo,
	}
}

func (t *Task) CreateTask(task *domain.Task) (*domain.Task, error) {
	return t.repo.Task.CreateTask(task)
}

func (t *Task) UpdateTaskOfUser(task *domain.Task) error {
	return t.repo.Task.UpdateTaskOfUser(task)
}

func (t *Task) DeleteUserTask(id int, user *domain.User) error {
	return t.repo.Task.DeleteUserTask(id, user)
}

func (t *Task) GetTaskListOfUser(userId int) ([]domain.Task, error) {
	return t.repo.Task.GetTaskListOfUser(userId)
}

func (t *Task) GetUserTaskByID(taskId int, user *domain.User) (*domain.Task, error) {
	return t.repo.Task.GetUserTaskByID(taskId, user)
}

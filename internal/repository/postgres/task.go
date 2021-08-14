package postgres

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mihail-1212/todo-project-backend/pkg/domain"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{
		db: db,
	}
}

func (t *TaskPostgres) CreateTask(task *domain.Task) (*domain.Task, error) {
	newTask := task

	query := fmt.Sprintf("INSERT INTO %s (name, userId, description, dateStart) VALUES ($1, $2, $3, $4) RETURNING id", taskTable)
	row := t.db.QueryRow(query, &task.Name, &task.UserId, &task.Description, &task.DateStart)
	if err := row.Scan(&newTask.Id); err != nil {
		if err.Error() == fmt.Sprintf(
			"pq: insert or update on table \"%s\" violates foreign key constraint \"%s\"", taskTable, userTable,
		) {
			return nil, errors.New("no user with this id")
		}
		return nil, err
	}

	return newTask, nil
}

func (t *TaskPostgres) UpdateTaskOfUser(task *domain.Task) error {
	var updatedTaskId int

	query := fmt.Sprintf("UPDATE %s SET name=$1, description=$2, dateStart=$3 WHERE id=$4 AND userId=$5 RETURNING id", taskTable)
	row := t.db.QueryRow(query, &task.Name, &task.Description, &task.DateStart, &task.Id, &task.UserId)
	if err := row.Scan(&updatedTaskId); err != nil {
		return err
	}

	return nil
}

func (t *TaskPostgres) DeleteUserTask(taskId int, user *domain.User) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND userId=$2", taskTable)
	_, err := t.db.Exec(query, taskId, user.Id)

	return err
}

func (t *TaskPostgres) GetTaskListOfUser(userId int) ([]domain.Task, error) {
	var tasks []domain.Task

	query := fmt.Sprintf("SELECT id, name, userId, description, dateStart FROM %s WHERE userId=$1", taskTable)

	rows, err := t.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task domain.Task
		err := rows.Scan(&task.Id, &task.Name, &task.UserId, &task.Description, &task.DateStart)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *TaskPostgres) GetUserTaskByID(taskId int, user *domain.User) (*domain.Task, error) {
	var task domain.Task

	query := fmt.Sprintf("SELECT id, name, userId, description, dateStart FROM %s WHERE id=$1 AND userId=$2", taskTable)

	row := t.db.QueryRow(query, taskId, user.Id)
	err := row.Scan(&task.Id, &task.Name, &task.UserId, &task.Description, &task.DateStart)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

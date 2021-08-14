package domain

import "time"

type Task struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	UserId      int       `json:"userid"`
	Description string    `json:"description"`
	DateStart   time.Time `json:"datestart"`
}

type CreateTask struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DateStart   time.Time `json:"datestart"`
}

type UpdateTask struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	DateStart   time.Time `json:"datestart"`
}

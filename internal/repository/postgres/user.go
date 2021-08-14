package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mihail-1212/todo-project-backend/pkg/auth/models"
	"github.com/mihail-1212/todo-project-backend/pkg/domain"
)

/*

	id serial NOT NULL PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	"password" VARCHAR ( 255 ) NOT NULL

*/

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (u *UserPostgres) CreateUser(user domain.User) (int, error) {
	return 0, nil
}

func (u *UserPostgres) IsUserExists(userId int) (bool, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE id=$1", userTable)
	row := u.db.QueryRow(query, userId)
	err := row.Scan()
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (u *UserPostgres) DeleteUser(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", userTable)
	_, err := u.db.Exec(query, id)

	return err
}

func (u *UserPostgres) Insert(user *models.User, generatePasswordHash func(string) string) error {
	// TODO: make insert
	return nil
}

func (u *UserPostgres) Get(username, password string, generatePasswordHash func(string) string) (*models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT username, \"password\" FROM %s WHERE username=$1 AND \"password\"=$2", userTable)
	row := u.db.QueryRow(query, username, generatePasswordHash(password))

	err := row.Scan(&user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserPostgres) GetUserByID(id int) (*domain.User, error) {
	var user domain.User

	query := fmt.Sprintf("SELECT id, username, \"password\" FROM %s WHERE id=$1", userTable)
	row := u.db.QueryRow(query, id)

	err := row.Scan(&user.Id, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserPostgres) GetUserByLogin(login string) (*domain.User, error) {
	var user domain.User

	query := fmt.Sprintf("SELECT id, username, \"password\" FROM %s WHERE username=$1", userTable)
	row := u.db.QueryRow(query, login)

	err := row.Scan(&user.Id, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

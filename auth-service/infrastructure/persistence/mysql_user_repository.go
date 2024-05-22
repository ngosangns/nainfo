package persistence

import (
	"auth-service/domain/model"
	"auth-service/domain/repository"
	"context"
	"database/sql"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) repository.UserRepository {
	return &MySQLUserRepository{db}
}

func (r *MySQLUserRepository) Save(user *model.User) error {
	tx, err := r.StartTransaction()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
		}
	}()

	return r.SaveWithTx(tx, user)
}

func (r *MySQLUserRepository) SaveWithTx(tx *sql.Tx, user *model.User) error {
	_, err := tx.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", user.Username, user.Password, user.Email)
	return err
}

func (r *MySQLUserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT id, username, password, email FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MySQLUserRepository) StartTransaction() (*sql.Tx, error) {
	return r.db.BeginTx(context.Background(), nil)
}

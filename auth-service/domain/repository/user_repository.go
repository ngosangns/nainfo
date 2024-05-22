package repository

import (
	"auth-service/domain/model"
	"database/sql"
)

type UserRepository interface {
	Save(user *model.User) error
	SaveWithTx(tx *sql.Tx, user *model.User) error
	FindByUsername(username string) (*model.User, error)
	StartTransaction() (*sql.Tx, error)
}

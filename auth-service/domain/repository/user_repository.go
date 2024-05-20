package repository

import "auth-service/domain/model"

type UserRepository interface {
	Save(user *model.User) error
	FindByUsername(username string) (*model.User, error)
}

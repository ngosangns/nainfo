package repository

import "profile-service/domain/model"

type ProfileRepository interface {
	Update(profile *model.Profile) error
	UpdateOrCreate(profile *model.Profile) error
	FindByUsername(username string) (*model.Profile, error)
}

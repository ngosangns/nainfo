package persistence

import (
	"database/sql"
	"profile-service/domain/model"
	"profile-service/domain/repository"
)

type MySQLProfileRepository struct {
	db *sql.DB
}

func NewMySQLProfileRepository(db *sql.DB) repository.ProfileRepository {
	return &MySQLProfileRepository{db}
}

func (r *MySQLProfileRepository) Update(profile *model.Profile) error {
	_, err := r.db.Exec("UPDATE profiles SET email = ? WHERE username = ?", profile.Email, profile.Username)
	return err
}

func (r *MySQLProfileRepository) FindByUsername(username string) (*model.Profile, error) {
	var profile model.Profile
	err := r.db.QueryRow("SELECT id, username, email FROM profiles WHERE username = ?", username).Scan(&profile.ID, &profile.Username, &profile.Email)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

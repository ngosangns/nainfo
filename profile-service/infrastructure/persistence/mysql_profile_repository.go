package persistence

import (
	"database/sql"
	"fmt"
	"profile-service/domain/model"
	"profile-service/domain/repository"
)

type MySQLProfileRepository struct {
	db *sql.DB
}

func NewMySQLProfileRepository(db *sql.DB) repository.ProfileRepository {
	return &MySQLProfileRepository{db}
}

func (r *MySQLProfileRepository) Create(profile *model.Profile) error {
	// Insert the new user into the database
	result, err := r.db.Exec("INSERT INTO profiles (username, email) VALUES (?, ?)", profile.Username, profile.Email)
	if err != nil {
		return err
	}

	// Get the ID of the newly created user
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Set the ID of the profile to the newly created ID
	profile.ID = uint64(id)

	return nil
}

func (r *MySQLProfileRepository) Update(profile *model.Profile) error {
	_, err := r.db.Exec("UPDATE profiles SET email = ? WHERE username = ?", profile.Email, profile.Username)
	return err
}

func (r *MySQLProfileRepository) UpdateOrCreate(profile *model.Profile) error {
	// Find the user by username
	result := r.db.QueryRow("SELECT id FROM profiles WHERE username = ?", profile.Username)
	fmt.Println(result)
	var id uint64
	err := result.Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	// If the user exists, update the email
	if id > 0 {
		return r.Update(profile)
	}

	// If the user doesn't exist, create a new user
	return r.Create(profile)
}

func (r *MySQLProfileRepository) FindByUsername(username string) (*model.Profile, error) {
	var profile model.Profile
	err := r.db.QueryRow("SELECT id, username, email FROM profiles WHERE username = ?", username).Scan(&profile.ID, &profile.Username, &profile.Email)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

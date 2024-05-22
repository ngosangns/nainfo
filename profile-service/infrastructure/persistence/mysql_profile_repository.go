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

func (r *MySQLProfileRepository) Create(profile *model.Profile) error {
	// Insert the new user into the database
	result, err := r.db.Exec("INSERT INTO profiles (username, name, description, email, address, facebook, linkedin, github) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		profile.Username,
		profile.Name,
		profile.Description,
		profile.Email,
		profile.Address,
		profile.Facebook,
		profile.LinkedIn, profile.GitHub)
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
	_, err := r.db.Exec("UPDATE profiles SET name = ?, description = ?, email = ?, address = ?, facebook = ?, linkedin = ?, github = ? WHERE username = ?",
		profile.Name,
		profile.Description,
		profile.Email,
		profile.Address,
		profile.Facebook,
		profile.LinkedIn,
		profile.GitHub, profile.Username)
	return err
}

func (r *MySQLProfileRepository) UpdateOrCreate(profile *model.Profile) error {
	// Find the user by username
	result := r.db.QueryRow("SELECT id FROM profiles WHERE username = ?", profile.Username)
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
	err := r.db.QueryRow("SELECT id, username, name, description, email, address, facebook, linkedin, github FROM profiles WHERE username = ?", username).Scan(&profile.ID, &profile.Username, &profile.Name, &profile.Description, &profile.Email, &profile.Address, &profile.Facebook, &profile.LinkedIn, &profile.GitHub)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

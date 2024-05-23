package service

import (
	"database/sql"
	"profile-service/domain/model"
	"profile-service/domain/repository"
	"profile-service/dto"
)

type ProfileService struct {
	profileRepo repository.ProfileRepository
}

func NewProfileService(repo repository.ProfileRepository) ProfileService {
	return ProfileService{profileRepo: repo}
}

func (s *ProfileService) UpdateProfile(req dto.UpdateProfileRequest) error {
	profile := model.Profile{
		Username: req.Username,
		Email:    req.Email,
	}

	return s.profileRepo.Update(&profile)
}

func (s *ProfileService) UpdateOrCreateProfile(req dto.UpdateProfileRequest) error {
	profile := model.Profile{
		Username:    req.Username,
		Email:       req.Email,
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		Facebook:    req.Facebook,
		LinkedIn:    req.LinkedIn,
		GitHub:      req.GitHub,
	}

	return s.profileRepo.UpdateOrCreate(&profile)
}

func (s *ProfileService) GetProfile(username string) (*model.Profile, error) {
	profile, err := s.profileRepo.FindByUsername(username)

	// Create if not exist
	if err != nil && err == sql.ErrNoRows {
		err = s.UpdateOrCreateProfile(dto.UpdateProfileRequest{Username: username})
		if err != nil {
			return nil, err
		}
		profile, err = s.profileRepo.FindByUsername(username)
	}

	if err != nil {
		return nil, err
	}

	return profile, nil
}

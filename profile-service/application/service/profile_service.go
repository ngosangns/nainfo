package service

import (
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

func (s *ProfileService) GetProfile(username string) (*model.Profile, error) {
	return s.profileRepo.FindByUsername(username)
}

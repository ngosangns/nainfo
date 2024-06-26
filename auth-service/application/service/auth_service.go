package service

import (
	"auth-service/domain/model"
	"auth-service/domain/repository"
	"auth-service/dto"
	"auth-service/infrastructure/grpc"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo      repository.UserRepository
	profileClient *grpc.ProfileClient
}

func NewAuthService(repo repository.UserRepository) AuthService {
	profileClient, err := grpc.NewProfileClient()
	if err != nil {
		panic(err)
	}

	return AuthService{userRepo: repo, profileClient: profileClient}
}

func (s *AuthService) Login(req dto.LoginRequest) (*model.User, error) {
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Register(req dto.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	tx, err := s.userRepo.StartTransaction()
	if err != nil {
		return err
	}
	transactionFailed := false
	defer func() {
		if transactionFailed {
			tx.Rollback()
		}
	}()

	err = s.userRepo.SaveWithTx(tx, &user)
	if err != nil {
		transactionFailed = true
		return err
	}

	// Register user in profile-service using gRPC
	// err = s.profileClient.UpdateOrCreateProfile(&proto.UpdateProfileRequest{
	// 	Username: req.Username,
	// 	Email:    req.Email,
	// })
	// if err != nil {
	// 	transactionFailed = true
	// 	return err
	// }

	// Commit transaction
	if err := tx.Commit(); err != nil {
		transactionFailed = true
		return err
	}

	return nil
}

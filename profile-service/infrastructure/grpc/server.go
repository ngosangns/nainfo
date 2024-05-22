package grpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"profile-service/application/service"
	"profile-service/domain/repository"
	"profile-service/dto"
	"shared/proto" // Generated from protobuf

	"google.golang.org/grpc"
)

type profileServer struct {
	proto.UnimplementedProfileServiceServer
	profileService service.ProfileService
}

func NewProfileServer(profileRepo repository.ProfileRepository) *profileServer {
	return &profileServer{
		profileService: service.NewProfileService(profileRepo),
	}
}

func (s *profileServer) GetProfile(ctx context.Context, req *proto.GetProfileRequest) (*proto.ProfileResponse, error) {
	profile, err := s.profileService.GetProfile(req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to get profile: %w", err)
	}

	return &proto.ProfileResponse{
		Username: profile.Username,
		Email:    profile.Email,
	}, nil
}

func (s *profileServer) UpdateProfile(ctx context.Context, req *proto.UpdateProfileRequest) (*proto.Empty, error) {
	err := s.profileService.UpdateOrCreateProfile(dto.UpdateProfileRequest{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update profile: %w", err)
	}

	return &proto.Empty{}, nil
}

func (s *profileServer) UpdateOrCreateProfile(ctx context.Context, req *proto.UpdateProfileRequest) (*proto.Empty, error) {
	err := s.profileService.UpdateOrCreateProfile(dto.UpdateProfileRequest{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update profile: %w", err)
	}

	return &proto.Empty{}, nil
}

func RunGRPCServer(profileRepo repository.ProfileRepository) error {
	lis, err := net.Listen("tcp", os.Getenv("PROFILE_SERVICE_GRPC_ADDRESS"))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	server := grpc.NewServer()
	profileServer := NewProfileServer(profileRepo)
	proto.RegisterProfileServiceServer(server, profileServer)

	if err := server.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

package grpc

import (
	"context"
	"fmt"
	"os"
	"time"

	"shared/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProfileClient struct {
	client proto.ProfileServiceClient
}

func NewProfileClient() (*ProfileClient, error) {
	conn, err := grpc.NewClient(os.Getenv("PROFILE_SERVICE_GRPC_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithIdleTimeout(time.Second))
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to connect to profile server: %w", err)
	}

	client := proto.NewProfileServiceClient(conn)
	return &ProfileClient{client: client}, nil
}

func (c *ProfileClient) GetProfile(username string) (*proto.ProfileResponse, error) {
	resp, err := c.client.GetProfile(context.Background(), &proto.GetProfileRequest{Username: username})
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to get profile: %w", err)
	}

	return resp, nil
}

func (c *ProfileClient) UpdateProfile(req *proto.UpdateProfileRequest) error {
	_, err := c.client.UpdateProfile(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to update profile: %w", err)
	}

	return nil
}

func (c *ProfileClient) UpdateOrCreateProfile(req *proto.UpdateProfileRequest) error {
	_, err := c.client.UpdateOrCreateProfile(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to update or create profile: %w", err)
	}

	return nil
}

package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/agbgru/auth/pkg/auth_v1"
)

const grpcPort = 50052

type server struct {
	desc.UnimplementedAuthServer
}

// CreateUser creates new user.
func (s *server) CreateUser(ctx context.Context, req *desc.CreateUserRequest,
) (*desc.CreateUserResponse, error) {
	_ = ctx

	log.Printf("CreateUser called with name: %s, email: %s, password: %s, password confirm: %s, role: %s\n",
		req.GetName(), req.GetEmail(), req.GetPassword(), req.GetPasswordConfirm(), req.GetRole().String())

	return &desc.CreateUserResponse{Id: 1}, nil
}

// GetUser retrieves a user by ID.
func (s *server) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	_ = ctx

	log.Printf("GetUser called with id: %d\n", req.GetId())

	return &desc.GetUserResponse{
		Id:        req.GetId(),
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		Role:      desc.Role_USER,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}, nil
}

// UpdateUser updates user details.
func (s *server) UpdateUser(ctx context.Context, req *desc.UpdateUserRequest) (*desc.UpdateUserResponse, error) {
	_ = ctx

	log.Printf("UpdateUser called with id: %d, name: %s, email: %s, role: %s\n",
		req.GetId(), req.GetName(), req.GetEmail(), req.GetRole().String())

	return &desc.UpdateUserResponse{}, nil
}

// DeleteUser deletes a user by ID.
func (s *server) DeleteUser(ctx context.Context, req *desc.DeleteUserRequest) (*desc.DeleteUserResponse, error) {
	_ = ctx

	log.Printf("DeleteUser called with id: %d\n", req.GetId())

	return &desc.DeleteUserResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

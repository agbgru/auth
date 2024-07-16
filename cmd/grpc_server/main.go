package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/agbgru/auth/pkg/auth_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedAuthServer
}

// CreateUser creates new user.
func (s *server) CreateUser(ctx context.Context, req *desc.CreateUserRequest,
) (*desc.CreateUserResponse, error) {
	_ = ctx

	log.Printf("CreateUser called with name: %s, email: %s, password: %s, password confirm: %s\n",
		req.Name, req.Email, req.Password, req.PasswordConfirm)

	return &desc.CreateUserResponse{Id: 1}, nil
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

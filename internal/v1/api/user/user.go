package user

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Psakine/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Server ...
type Server struct {
	user_v1.UnimplementedUserV1Server
}

// NewUserServer ...
func NewUserServer() *Server {
	return &Server{}
}

// Get ...
func (s *Server) Get(_ context.Context, _ *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	return &user_v1.GetResponse{
		Id:        gofakeit.Int64(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      user_v1.Role_USER,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}, nil
}

// Create ...
func (s *Server) Create(_ context.Context, _ *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	return &user_v1.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

// Update ...
func (s *Server) Update(_ context.Context, _ *user_v1.UpdateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// Delete ...
func (s *Server) Delete(_ context.Context, _ *user_v1.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

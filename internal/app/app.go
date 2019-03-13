package app

import (
	"context"

	pb "github.com/mjlaufer/go-grpc-service/api/proto"
)

// Service implements the methods defined in the protobuf definition and the generated protobuf interface
type Service struct {
	users []*pb.User
}

// Create user
func (s *Service) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	s.users = append(s.users, req.User)
	return &pb.CreateResponse{Id: req.User.Id}, nil
}

// List users
func (s *Service) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	users := s.users
	return &pb.ListResponse{Users: users}, nil
}

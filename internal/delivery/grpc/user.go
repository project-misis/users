package grpc

import (
	"context"

	"github.com/project-misis/users_proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedCrudServer
}

var _ pb.CrudClient = (*Server)(nil)

func (s *Server) CreateUser(context.Context, *pb.UserPost) (*pb.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func (s *Server) DeleteUserById(context.Context, *pb.UserDelete) (*pb.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserById not implemented")
}

func (s *Server) UpdateUserById(context.Context, *pb.UserUpdate) (*pb.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserById not implemented")
}

func (s *Server) GetUserById(context.Context, *pb.UserGet) (*pb.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}

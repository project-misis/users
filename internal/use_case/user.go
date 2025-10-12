package usecase

import (
	"context"

	"echo-template/internal/infrastructure/repository"

	"github.com/project-misis/users_proto/pb"
)

type UserService struct {
	clientRepo *repository.UserRepository
}

func NewClientService(repo *repository.UserRepository) *UserService {
	return &UserService{clientRepo: repo}
}

func (s *UserService) CreateUser(up *pb.UserPost) {
}

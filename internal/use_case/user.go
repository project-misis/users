// Package usecase
package usecase

import (
	"context"

	"echo-template/db"
	"echo-template/internal/infrastructure/repository"

	"github.com/google/uuid"
	"github.com/project-misis/users_proto/pb"
)

type UserService struct {
	r *repository.UserRepository
}

func NewClientService(repo *repository.UserRepository) *UserService {
	return &UserService{r: repo}
}

func (s *UserService) GetUserByID(ctx context.Context, p *pb.UserGet) (*pb.User, error) {
	uid, err := uuid.Parse(p.GetId())
	if err != nil {
		return nil, err
	}
	u, err := s.r.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:        u.ID.String(),
		Username:  u.Username,
		FirstName: u.Firstname,
		Course:    u.Course,
		Faculty:   u.Faculty,
	}, nil
}

func (s *UserService) DeleteUserByID(ctx context.Context, p *pb.UserDelete) error {
	uid, err := uuid.Parse(p.GetId())
	if err != nil {
		return err
	}
	err = s.r.DeleteUser(ctx, uid)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUserByID(ctx context.Context, p *pb.UserUpdate) (*pb.User, error) {
	uid, err := uuid.Parse(p.GetId())
	if err != nil {
		return nil, err
	}
	params := db.UpdateUserParams{
		ID:        uid,
		Username:  p.Username,
		Firstname: p.FirstName,
		Faculty:   p.Faculty,
		Course:    p.Course,
	}
	u, err := s.r.UpdateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:        u.ID.String(),
		Username:  u.Username,
		FirstName: u.Firstname,
		Course:    u.Course,
		Faculty:   u.Faculty,
	}, nil
}

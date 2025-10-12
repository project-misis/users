package repository

import (
	"context"

	"echo-template/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewClientRepository(db *pgxpool.Pool) *UserRepository {
	if db == nil {
		panic("Database connection is nil in repository")
	}
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, p db.CreateUserParams) error {
	q := db.New(r.db)
	return q.CreateUser(ctx, p)
}

func (r *UserRepository) GetUser(ctx context.Context, id uuid.UUID) (*db.User, error) {
	q := db.New(r.db)
	user, err := q.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	q := db.New(r.db)
	return q.DeleteUser(ctx, id)
}

func (r *UserRepository) UpdateUser(ctx context.Context, p db.UpdateUserParams) (*db.User, error) {
	q := db.New(r.db)
	user, err := q.UpdateUser(ctx, p)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

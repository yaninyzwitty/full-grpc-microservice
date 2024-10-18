package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yaninyzwitty/grpc-microservice-postgres/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	CreatePost(ctx context.Context, post models.Post) (*models.Post, error)
	LikeComment(ctx context.Context, comment models.Comment) (*models.Comment, error)
	LikePost(ctx context.Context, post models.Post) (*models.Post, error)
	CreateComment(ctx context.Context, comment models.Comment) (*models.Comment, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db}
}
func (r *userRepository) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	return &models.User{}, nil
}
func (r *userRepository) CreatePost(ctx context.Context, post models.Post) (*models.Post, error) {
	return &models.Post{}, nil
}
func (r *userRepository) CreateComment(ctx context.Context, comment models.Comment) (*models.Comment, error) {
	return &models.Comment{}, nil
}
func (r *userRepository) LikePost(ctx context.Context, post models.Post) (*models.Post, error) {
	return &models.Post{}, nil
}
func (r *userRepository) LikeComment(ctx context.Context, comment models.Comment) (*models.Comment, error) {
	return &models.Comment{}, nil
}

package service

import (
	"context"

	"github.com/yaninyzwitty/grpc-microservice-postgres/internal/models"
	"github.com/yaninyzwitty/grpc-microservice-postgres/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	CreatePost(ctx context.Context, post models.Post) (*models.Post, error)
	LikeComment(ctx context.Context, comment models.Comment) (*models.Comment, error)
	LikePost(ctx context.Context, post models.Post) (*models.Post, error)
	CreateComment(ctx context.Context, comment models.Comment) (*models.Comment, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) UserService {
	return &userService{repo: *repo}
}

func (s *userService) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	return s.repo.CreateUser(ctx, user)
}
func (s *userService) CreateComment(ctx context.Context, comment models.Comment) (*models.Comment, error) {
	return s.repo.CreateComment(ctx, comment)
}
func (s *userService) CreatePost(ctx context.Context, post models.Post) (*models.Post, error) {
	return s.repo.CreatePost(ctx, post)
}
func (s *userService) LikeComment(ctx context.Context, comment models.Comment) (*models.Comment, error) {
	return s.repo.LikeComment(ctx, comment)
}
func (s *userService) LikePost(ctx context.Context, post models.Post) (*models.Post, error) {
	return s.repo.LikePost(ctx, post)
}

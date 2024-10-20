package controller

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yaninyzwitty/grpc-microservice-postgres/internal/models"
	"github.com/yaninyzwitty/grpc-microservice-postgres/internal/service"
	"github.com/yaninyzwitty/grpc-microservice-postgres/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserController struct {
	service service.UserService
	pb.UnimplementedPhotoSharingServiceServer
}

func NewUserController(svc service.UserService) *UserController {
	return &UserController{service: svc}
}

func (c *UserController) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	userId := uuid.New().String()
	user := models.User{
		ID:        userId,
		Username:  req.Username,
		Name:      req.Name,
		Email:     req.Email,
		Bio:       req.Bio,
		ImageUrl:  req.ImageUrl,
		CreatedAt: time.Now(),
	}

	createdUser, err := c.service.CreateUser(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create user %v", err)
	}
	return &pb.User{
		Id:        createdUser.ID,
		Username:  createdUser.Username,
		Name:      createdUser.Name,
		Email:     createdUser.Email,
		Bio:       createdUser.Bio,
		ImageUrl:  createdUser.ImageUrl,
		CreatedAt: timestamppb.New(createdUser.CreatedAt),
	}, nil
}
func (c *UserController) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
	postId := uuid.New().String()
	post := models.Post{
		Id:        postId,
		Content:   req.Content,
		AuthorId:  req.AuthorId,
		Likes:     0,
		CreatedAt: time.Now(),
	}

	createdPost, err := c.service.CreatePost(ctx, post)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create post %v", err)
	}
	return &pb.Post{
		Id:        createdPost.Id,
		Content:   createdPost.Content,
		AuthorId:  createdPost.AuthorId,
		Likes:     int32(createdPost.Likes),
		CreatedAt: timestamppb.New(createdPost.CreatedAt),
	}, nil
}
func (c *UserController) CreateComment(ctx context.Context, req *pb.CreateCommentInput) (*pb.Comment, error) {
	commentId := uuid.New().String()
	comment := models.Comment{
		ID:        commentId,
		Content:   req.Content,
		PostId:    req.PostId,
		UserId:    req.UserId,
		CreatedAt: time.Now(),
		Likes:     0,
	}

	createdComment, err := c.service.CreateComment(ctx, comment)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create comment %v", err)
	}
	return &pb.Comment{
		Id:        createdComment.ID,
		Content:   createdComment.Content,
		PostId:    createdComment.PostId,
		UserId:    createdComment.UserId,
		Likes:     int32(createdComment.Likes),
		CreatedAt: timestamppb.New(comment.CreatedAt),
	}, nil
}

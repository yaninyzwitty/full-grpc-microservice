package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yaninyzwitty/grpc-microservice-postgres/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	CreatePost(ctx context.Context, post models.Post) (*models.Post, error)
	LikeComment(ctx context.Context, commentId string) (*models.Comment, error)
	LikePost(ctx context.Context, postId string) (*models.Post, error)
	CreateComment(ctx context.Context, comment models.Comment) (*models.Comment, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db}
}
func (r *userRepository) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	query := `INSERT INTO users (id, username, name, email, bio, image_url, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	res, err := r.db.Exec(ctx, query, user.ID, user.Username, user.Name, user.Email, user.Bio, user.ImageUrl, user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	rowsAffected := res.RowsAffected()

	if rowsAffected == 0 {
		return nil, fmt.Errorf("no rows were inserted")
	}
	return &user, nil

}
func (r *userRepository) CreatePost(ctx context.Context, post models.Post) (*models.Post, error) {
	query := `INSERT INTO posts (id, content, author_id, likes, created_at) VALUES($1, $2, $3, $4, $5)`

	res, err := r.db.Exec(ctx, query, post.Id, post.Content, post.AuthorId, post.Likes, post.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to insert post: %w", err)
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no rows were inserted")
	}

	return &post, nil
}
func (r *userRepository) CreateComment(ctx context.Context, comment models.Comment) (*models.Comment, error) {
	query := `INSERT INTO comments (id, content, post_id, user_id, likes, created_at) VALUES($1, $2, $3, $4, $5, $6)`

	res, err := r.db.Exec(ctx, query, comment.ID, comment.Content, comment.PostId, comment.UserId, comment.Likes, comment.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to insert comment: %w", err)
	}
	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no rows were inserted")
	}

	return &comment, nil
}
func (r *userRepository) LikePost(ctx context.Context, postId string) (*models.Post, error) {
	query := `UPDATE posts SET likes = likes + 1 WHERE id = $1 RETURNING id, content,  author_id, likes, created_at`
	var updatedPost models.Post
	err := r.db.QueryRow(ctx, query, postId).Scan(&updatedPost.Id, &updatedPost.Content, &updatedPost.AuthorId, &updatedPost.Likes, &updatedPost.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("post with id %s not found", postId)
		}
		return nil, fmt.Errorf("failed to like post with id: %s: %w", postId, err)
	}
	return &updatedPost, nil
}
func (r *userRepository) LikeComment(ctx context.Context, commentId string) (*models.Comment, error) {
	query := `UPDATE comments SET likes = likes + 1 WHERE id = $1 RETURNING id, content, post_id, user_id, likes, created_at`
	var updatedComment models.Comment
	err := r.db.QueryRow(ctx, query, commentId).Scan(&updatedComment.ID, &updatedComment.Content, &updatedComment.UserId, &updatedComment.Likes, &updatedComment.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("comment with id %s not found", commentId)
		}
		return nil, fmt.Errorf("failed to like comment with id: %s: %w", commentId, err)
	}
	return &updatedComment, nil
}

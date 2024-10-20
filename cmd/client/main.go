package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/yaninyzwitty/grpc-microservice-postgres/pb"
	"github.com/yaninyzwitty/grpc-microservice-postgres/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var cfg pkg.Config

	file, err := os.Open("config.yaml")
	if err != nil {
		slog.Error("Failed to open  config.yaml file")
		os.Exit(1)
	}

	defer file.Close()
	if err := cfg.LoadConfig(file); err != nil {
		slog.Error("Error loading config.yaml", "error", err)
		os.Exit(1)
	}

	address := fmt.Sprintf(":%d", 50051)

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("failed to create a grpc conn", "error", err)
	}
	defer conn.Close()
	client := pb.NewPhotoSharingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	createUserResult, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		Username: "kALI witty",
		Name:     "Ian Mwangi Munyiri (kw)",
		Email:    "kaliwitty@outlook.com",
		Bio:      "Witty is a brilliant name",
		ImageUrl: "https://images.unsplash.com/photo-1719518870616-8deacda7e18b?w=800&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHx0b3BpYy1mZWVkfDQwfHRvd0paRnNrcEdnfHxlbnwwfHx8fHw%3D",
	})
	if err != nil {
		slog.Error("Failed to create user", "error", err)
		return
	}

	slog.Info("User created: ", "res", createUserResult.Id)

	createPostResult, err := client.CreatePost(ctx, &pb.CreatePostRequest{
		Content:  "It's nice having chesticles",
		AuthorId: createUserResult.Id,
	})

	if err != nil {
		slog.Error("Failed to create post", "error", err)
		return
	}

	slog.Info("Created post: ", "res", createPostResult.Id)

	createCommentResult, err := client.CreateComment(ctx, &pb.CreateCommentInput{
		Content: "Yoh! I like that DAMN!💦",
		PostId:  createPostResult.Id,
		UserId:  createPostResult.AuthorId,
	})

	if err != nil {
		slog.Error("Failed to create comment", "error", err)
		return
	}
	slog.Info("Created comment: ", "res", createCommentResult.Id)

}

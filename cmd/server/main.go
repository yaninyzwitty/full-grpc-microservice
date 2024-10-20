package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yaninyzwitty/grpc-microservice-postgres/internal/controller"
	"github.com/yaninyzwitty/grpc-microservice-postgres/internal/database"
	"github.com/yaninyzwitty/grpc-microservice-postgres/internal/repository"
	"github.com/yaninyzwitty/grpc-microservice-postgres/internal/service"
	"github.com/yaninyzwitty/grpc-microservice-postgres/pb"
	"github.com/yaninyzwitty/grpc-microservice-postgres/pkg"
	"google.golang.org/grpc"
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
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Database.Timeout*time.Second)
	defer cancel()

	dbConfig := &database.DatabaseConfig{
		Username: cfg.Database.Username,
		Password: cfg.Database.Password,
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		DBName:   cfg.Database.DBName,
		SSLMode:  cfg.Database.SSLMode,
	}

	db, err := database.NewDatabaseConnection(ctx, dbConfig)
	if err != nil {
		slog.Error("Failed to create a database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	err = database.PingDatabase(ctx, db)
	if err != nil {
		slog.Error("Failed to ping database", "error", err)
	}

	// setup grpc server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Database.Port))
	if err != nil {
		slog.Error("Failed to listen", "error", err)
		os.Exit(1)
	}
	// set up injects 💉
	userRepo := repository.NewRepository(db)
	userService := service.NewUserService(&userRepo)
	userController := controller.NewUserController(userService)

	server := grpc.NewServer()
	pb.RegisterPhotoSharingServiceServer(server, userController)
	slog.Info("Server is listening", "address", listener.Addr().String())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-sigChan
		slog.Info("Received shutdown signal", "signal", sig)
		slog.Info("Shutting down gRPC server...")

		// Gracefully stop the gRPC server
		server.GracefulStop()
		cancel()

		slog.Info("gRPC server has been stopped gracefully")
	}()

	slog.Info("Starting gRPC server", "port", cfg.Server.PORT)
	if err := server.Serve(listener); err != nil {
		slog.Error("gRPC server encountered an error while serving", "error", err)
		os.Exit(1)
	}
}

package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/yaninyzwitty/grpc-microservice-postgres/internal/database"
	"github.com/yaninyzwitty/grpc-microservice-postgres/pkg"
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
	ctx, cancel := context.WithTimeout(context.Background(), 65*time.Second)
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
	defer db.Close(ctx)

	err = database.PingDatabase(ctx, db)
	if err != nil {
		slog.Error("Failed to ping database")
	}

}

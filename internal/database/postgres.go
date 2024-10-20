package database

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
	SSLMode  string
}

func NewDatabaseConnection(ctx context.Context, config *DatabaseConfig) (*pgxpool.Pool, error) {
	if config.Username == "" || config.Password == "" || config.Host == "" || config.DBName == "" || config.Port == 0 {
		return &pgxpool.Pool{}, errors.New("missing required database configuration")
	}

	// connectionUrl := fmt.Sprintf(
	// 	"postgres://%s:%s@%s:%d/%s?sslmode=%s",
	// 	config.Username,
	// 	config.Password,
	// 	config.Host,
	// 	config.Port,
	// 	config.DBName,
	// 	config.SSLMode,
	// )
	// "postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable"
	connectionPool, err := pgxpool.New(ctx, "postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable")
	if err != nil {
		slog.Error("failed to connect to db", "error", err)
	}
	return connectionPool, nil

}

func PingDatabase(ctx context.Context, conn *pgxpool.Pool) error {
	const timeout = 60 * time.Second
	const pingInterval = 1 * time.Second

	endTime := time.Now().Add(timeout)

	for time.Now().Before(endTime) {
		if err := conn.Ping(ctx); err != nil {
			slog.Error("Failed to ping database", "error", err)
			time.Sleep(pingInterval)
			continue
		}

		slog.Info("Successfully pinged database")
		return nil
	}

	slog.Info("Completed 60 seconds of pinging the database without success")
	return fmt.Errorf("database not reachable within timeout period")
}

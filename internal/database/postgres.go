package database

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
	SSLMode  string
}

func NewDatabaseConnection(ctx context.Context, config *DatabaseConfig) (*pgx.Conn, error) {
	if config.Username == "" || config.Password == "" || config.Host == "" || config.DBName == "" || config.Port == 0 {
		return &pgx.Conn{}, errors.New("missing required database configuration")
	}

	connectionUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.Username, config.Password, config.Host, config.Port, config.DBName, config.SSLMode,
	)
	conn, err := pgx.Connect(ctx, connectionUrl)
	if err != nil {
		slog.Error("failed to connect to db", "error", err)
	}
	return conn, nil

}

func PingDatabase(ctx context.Context, conn *pgx.Conn) error {
	timeout := 60 * time.Second
	endTime := time.Now().Add(timeout)

	for time.Now().Before(endTime) {
		err := conn.Ping(ctx)
		if err != nil {
			slog.Error("failed to ping database", "error", err)
			return err
		}

		slog.Info("successfully pinged database")
		time.Sleep(1 * time.Second) // Wait 1 second before next ping
	}

	slog.Info("completed 60 seconds of pinging the database")
	return nil
}

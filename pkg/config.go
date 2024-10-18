package pkg

import (
	"io"
	"log/slog"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server   Server `yaml:"server"`
	Database DB     `yaml:"database"`
}

type Server struct {
	PORT int `yaml:"port"`
}

type DB struct {
	Username string        `yaml:"username"`
	Password string        `yaml:"password"`
	Host     string        `yaml:"host"`
	Port     int           `yaml:"port"`
	DBName   string        `yaml:"db_name"`
	SSLMode  string        `yaml:"ssl_mode"`
	Timeout  time.Duration `yaml:"timeout"`
}

func (c *Config) LoadConfig(file io.Reader) error {
	data, err := io.ReadAll(file)
	if err != nil {
		slog.Error("Failed to read file", "error", err)
		return err
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		slog.Error("Failed to unmarshal file data", "error", err)
		return err
	}

	return nil

}

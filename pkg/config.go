package pkg

import (
	"io"
	"log/slog"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server   Server `yaml:"server"`
	Database DB     `yaml:"server"`
}

type Server struct {
	PORT int `yaml:"port"`
}

type DB struct {
	DATABASE_URL string `yaml:"database_url"`
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

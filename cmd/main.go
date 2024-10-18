package main

import (
	"log/slog"
	"os"

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

}

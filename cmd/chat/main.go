package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"

	"chat-service/internal/app"
	"chat-service/internal/config"
	"chat-service/internal/lib/logger"
)

// @title		Swagger Chat API
// @version		1.0
// @description	Chat service

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	if err := godotenv.Load(); err != nil {
		slog.Warn("failed load env", "error", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_PATH is not set")
	}

	cfg, err := config.New(configPath)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	chatApp := app.App{
		Config: cfg,
		Logger: logger.New(cfg.Env, cfg.LogLevel),
	}
	chatApp.MustLoad()
}

package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"chat-service/internal/config"
	"chat-service/internal/http/handler"
	"chat-service/internal/repository/postgres"
	"chat-service/internal/service"
)

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

	repo := postgres.New(cfg)
	repo.Migrate()

	svc := service.New(repo)
	h := handler.New(svc)

	mux := http.NewServeMux()
	h.Router(mux)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

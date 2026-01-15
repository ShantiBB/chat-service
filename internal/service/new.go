package service

import (
	"context"

	"chat-service/internal/repository/models"
)

type Repository interface {
	ChatRepository
}

type ChatRepository interface {
	InsertChat(ctx context.Context, chat *models.Chat) error
	SelectChat(ctx context.Context, id uint) (models.Chat, error)
	DeleteChat(ctx context.Context, id uint) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

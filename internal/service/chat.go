package service

import (
	"context"

	"chat-service/internal/repository/models"
)

func (s *Service) CreateChat(ctx context.Context, chat *models.Chat) error {
	return s.repo.InsertChat(ctx, chat)
}

func (s *Service) GetChatByID(ctx context.Context, id uint, limit int) (models.Chat, error) {
	chat, err := s.repo.SelectChat(ctx, id)
	if err != nil {
		return models.Chat{}, err
	}

	chat.Messages, err = s.repo.SelectMessages(ctx, id, limit)
	if err != nil {
		return models.Chat{}, err
	}

	return chat, nil
}

func (s *Service) DeleteChatByID(ctx context.Context, id uint) error {
	return s.repo.DeleteChat(ctx, id)
}

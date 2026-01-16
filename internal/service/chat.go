package service

import (
	"context"
	"strings"

	"chat-service/internal/repository/models"
)

func (s *Service) CreateChat(ctx context.Context, chat *models.Chat) error {
	if err := s.repo.InsertChat(ctx, chat); err != nil {
		return err
	}

	chat.Title = strings.TrimSpace(chat.Title)
	return nil
}

func (s *Service) GetChatByID(ctx context.Context, id uint, limit int) (*models.Chat, error) {
	chat, err := s.repo.SelectChat(ctx, id)
	if err != nil {
		return nil, err
	}

	chat.Messages, err = s.repo.SelectMessages(ctx, id, limit)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (s *Service) DeleteChatByID(ctx context.Context, id uint) error {
	if err := s.repo.DeleteChat(ctx, id); err != nil {
		return err
	}

	return nil
}

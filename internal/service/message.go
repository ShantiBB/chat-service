package service

import (
	"context"
	"strings"

	"chat-service/internal/repository/models"
)

func (s *Service) CreateMessage(ctx context.Context, msg *models.Message) error {
	msg.Text = strings.TrimSpace(msg.Text)

	if err := s.repo.InsertMessage(ctx, msg); err != nil {
		return err
	}

	return nil
}

package service

import (
	"context"

	"chat-service/internal/repository/models"
)

func (s *Service) CreateMessage(ctx context.Context, msg *models.Message) error {
	return s.repo.InsertMessage(ctx, msg)
}

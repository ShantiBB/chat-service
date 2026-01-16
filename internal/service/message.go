package service

import (
	"context"
	"errors"
	"log/slog"
	"strings"

	"chat-service/internal/lib/utils/consts"
	"chat-service/internal/repository/models"
)

func (s *Service) CreateMessage(ctx context.Context, msg *models.Message) error {
	msg.Text = strings.TrimSpace(msg.Text)

	if err := s.repo.InsertMessage(ctx, msg); err != nil {
		if !errors.Is(err, consts.ErrChatNotFound) {
			slog.Error("failed create message", "error", err)
		}
		return err
	}

	return nil
}

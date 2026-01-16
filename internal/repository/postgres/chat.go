package postgres

import (
	"context"
	"errors"
	"log/slog"

	"gorm.io/gorm"

	"chat-service/internal/lib/utils/consts"
	"chat-service/internal/repository/models"
)

func (r *Repository) InsertChat(ctx context.Context, chat *models.Chat) error {
	if err := r.db.WithContext(ctx).Create(chat).Error; err != nil {
		slog.Error("failed to create chat", "error", err)
		return err
	}

	return nil
}

func (r *Repository) SelectChat(ctx context.Context, id uint) (*models.Chat, error) {
	var chat models.Chat
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&chat).Error; err != nil {
		slog.Error("failed to select chat", "error", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, consts.ErrChatNotFound
		}

		return nil, err
	}

	return &chat, nil
}

func (r *Repository) DeleteChat(ctx context.Context, id uint) error {
	res := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Chat{})

	if res.Error != nil {
		slog.Error("failed to delete chat", "error", res.Error)
		return res.Error
	}

	if res.RowsAffected == 0 {
		return consts.ErrChatNotFound
	}

	return nil
}

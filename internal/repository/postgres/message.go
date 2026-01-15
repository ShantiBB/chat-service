package postgres

import (
	"context"

	"chat-service/internal/repository/models"
)

func (r *Repository) InsertMessage(ctx context.Context, message *models.Message) error {
	return r.db.WithContext(ctx).Create(message).Error
}

func (r *Repository) SelectMessages(ctx context.Context, chatID uint, limit int) ([]models.Message, error) {
	var messages []models.Message
	err := r.db.
		WithContext(ctx).
		Where("chat_id = ?", chatID).
		Order("created_at DESC, id DESC").
		Limit(limit).
		Find(&messages).
		Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}

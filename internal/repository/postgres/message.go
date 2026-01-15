package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"

	"chat-service/internal/repository/models"
	"chat-service/internal/utils/consts"
)

func (r *Repository) InsertMessage(ctx context.Context, message *models.Message) error {
	if err := r.db.WithContext(ctx).Create(message).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == consts.ForeignKeyViolates {
				return consts.ChatNotFound
			}
		}
	}

	return nil
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

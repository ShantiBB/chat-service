package postgres

import (
	"context"
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5/pgconn"

	consts2 "chat-service/internal/lib/utils/consts"
	"chat-service/internal/repository/models"
)

func (r *Repository) InsertMessage(ctx context.Context, message *models.Message) error {
	if err := r.db.WithContext(ctx).Create(message).Error; err != nil {
		slog.Error("failed to insert message", "error", err)

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == consts2.ForeignKeyViolates {
				return consts2.ErrChatNotFound
			}
		}

		return err
	}

	return nil
}

func (r *Repository) SelectMessages(ctx context.Context, chatID uint, limit int) ([]*models.Message, error) {
	var messages []*models.Message
	err := r.db.
		WithContext(ctx).
		Where("chat_id = ?", chatID).
		Order("created_at DESC, id DESC").
		Limit(limit).
		Find(&messages).
		Error

	if err != nil {
		slog.Error("failed to select messages", "error", err)
		return nil, err
	}

	return messages, nil
}

package postgres

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"chat-service/internal/repository/consts"
	"chat-service/internal/repository/models"
)

func (r *Repository) InsertChat(ctx context.Context, chat *models.Chat) error {
	return gorm.G[models.Chat](r.db).Create(ctx, chat)
}

func (r *Repository) SelectChat(ctx context.Context, id uint) (models.Chat, error) {
	chat, err := gorm.G[models.Chat](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Chat{}, consts.ChatNotFound
		}
		return models.Chat{}, err
	}

	return chat, nil
}

func (r *Repository) DeleteChat(ctx context.Context, id uint) error {
	rowsAffected, err := gorm.G[models.Chat](r.db).Where("id = ?", id).Delete(ctx)

	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return consts.FailedDeleteChat
	}

	return nil
}

package postgres

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"chat-service/internal/repository/models"
	"chat-service/internal/utils/consts"
)

func (r *Repository) InsertChat(ctx context.Context, chat *models.Chat) error {
	return r.db.WithContext(ctx).Create(chat).Error
}

func (r *Repository) SelectChat(ctx context.Context, id uint) (models.Chat, error) {
	var chat models.Chat
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&chat).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Chat{}, consts.ChatNotFound
		}

		return models.Chat{}, err
	}

	return chat, nil
}

func (r *Repository) DeleteChat(ctx context.Context, id uint) error {
	res := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Chat{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return consts.ChatNotFound
	}

	return nil
}

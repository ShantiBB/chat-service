package mappers

import (
	"chat-service/internal/http/dto/response"
	"chat-service/internal/repository/models"
)

func ChatToResponse(chat models.Chat) response.Chat {
	return response.Chat{
		ID:        chat.ID,
		Title:     chat.Title,
		CreatedAt: chat.CreatedAt,
	}
}

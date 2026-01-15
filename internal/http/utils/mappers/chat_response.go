package mappers

import (
	"chat-service/internal/http/dto/response"
	"chat-service/internal/repository/models"
)

func CreateChatToResponse(chat models.Chat) response.CreateChat {
	return response.CreateChat{
		ID:        chat.ID,
		Title:     chat.Title,
		CreatedAt: chat.CreatedAt,
	}
}

func ChatToResponse(chat models.Chat) response.Chat {
	return response.Chat{
		ID:        chat.ID,
		Title:     chat.Title,
		Messages:  MessagesToResponse(chat.Messages),
		CreatedAt: chat.CreatedAt,
	}
}

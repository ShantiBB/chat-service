package fixtures

import (
	"chat-service/internal/http/dto/request"
	"chat-service/internal/repository/models"
)

func GetCreateChatRequest(title string) request.CreateChat {
	return request.CreateChat{
		Title: title,
	}
}

func GetTestChat(id uint, title string) *models.Chat {
	return &models.Chat{
		ID:    id,
		Title: title,
	}
}

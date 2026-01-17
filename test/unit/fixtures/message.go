package fixtures

import (
	"chat-service/internal/http/dto/request"
	"chat-service/internal/repository/models"
)

func GetCreateMessageRequest(text string) request.CreateMessage {
	return request.CreateMessage{
		Text: text,
	}
}

func GetTestMessage(id, chatID uint, text string) *models.Message {
	return &models.Message{
		ID:     id,
		ChatID: chatID,
		Text:   text,
	}
}

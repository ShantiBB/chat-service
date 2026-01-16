package mappers

import (
	"chat-service/internal/http/dto/request"
	"chat-service/internal/repository/models"
)

func CreateMessageToModel(req *request.CreateMessage, chatID uint) *models.Message {
	return &models.Message{
		Text:   req.Text,
		ChatID: chatID,
	}
}

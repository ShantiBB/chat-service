package mappers

import (
	"chat-service/internal/http/dto/request"
	"chat-service/internal/repository/models"
)

func CreateChatToModel(req *request.CreateChat) *models.Chat {
	return &models.Chat{
		Title: req.Title,
	}
}

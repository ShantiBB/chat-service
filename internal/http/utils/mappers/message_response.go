package mappers

import (
	"chat-service/internal/http/dto/response"
	"chat-service/internal/repository/models"
)

func MessageToResponse(msg *models.Message) *response.Message {
	return &response.Message{
		ID:        msg.ID,
		Text:      msg.Text,
		CreatedAt: msg.CreatedAt,
	}
}

func MessagesToResponse(msgs []*models.Message) []*response.Message {
	msgsResp := make([]*response.Message, len(msgs))
	for i, msg := range msgs {
		msgsResp[i] = MessageToResponse(msg)
	}

	return msgsResp
}

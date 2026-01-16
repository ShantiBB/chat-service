package handler

import (
	"context"
	"net/http"

	"chat-service/internal/config"
	"chat-service/internal/repository/models"
)

type ChatService interface {
	CreateChat(ctx context.Context, chat *models.Chat) error
	GetChatByID(ctx context.Context, id uint, limit int) (*models.Chat, error)
	DeleteChatByID(ctx context.Context, id uint) error
}

type MessageService interface {
	CreateMessage(ctx context.Context, msg *models.Message) error
}

type Service interface {
	ChatService
	MessageService
}
type Handler struct {
	svc Service
	cfg *config.Config
}

func New(svc Service, cfg *config.Config) *Handler {
	return &Handler{
		svc: svc,
		cfg: cfg,
	}
}

func (h *Handler) Router(mux *http.ServeMux) {
	mux.HandleFunc("POST /chats", h.CreateChat)
	mux.HandleFunc("GET /chats/{id}", h.GetChatByID)
	mux.HandleFunc("DELETE /chats/{id}", h.DeleteChatByID)

	mux.HandleFunc("POST /chats/{chatID}/messages", h.CreateMessage)
}

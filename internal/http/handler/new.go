package handler

import (
	"context"
	"net/http"

	"chat-service/internal/repository/models"
)

type ChatService interface {
	CreateChat(ctx context.Context, chat *models.Chat) error
	GetChatByID(ctx context.Context, id uint) (models.Chat, error)
	DeleteChatByID(ctx context.Context, id uint) error
}

type Handler struct {
	svc ChatService
}

func New(svc ChatService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Router(mux *http.ServeMux) {
	mux.HandleFunc("POST /chats", h.CreateChat)
	mux.HandleFunc("GET /chats/{id}", h.GetChatByID)
	mux.HandleFunc("DELETE /chats/{id}", h.DeleteChatByID)
}

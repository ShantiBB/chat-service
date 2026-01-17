package handler

import (
	"context"
	"net/http"

	httpswagger "github.com/swaggo/http-swagger"

	_ "chat-service/docs"
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
	Svc Service
	Cfg *config.Config
}

func New(svc Service, cfg *config.Config) *Handler {
	return &Handler{
		Svc: svc,
		Cfg: cfg,
	}
}

func (h *Handler) Router(mux *http.ServeMux) {
	api := http.NewServeMux()

	api.HandleFunc("POST /chats", h.CreateChat)
	api.HandleFunc("GET /chats/{id}", h.GetChatByID)
	api.HandleFunc("DELETE /chats/{id}", h.DeleteChatByID)
	api.HandleFunc("POST /chats/{chatID}/messages", h.CreateMessage)

	mux.Handle("GET /api/v1/swagger/", httpswagger.WrapHandler)
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", api))
}

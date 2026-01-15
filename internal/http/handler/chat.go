package handler

import (
	"context"
	"net/http"
	"time"

	"chat-service/internal/http/dto/request"
	"chat-service/internal/http/utils/helpers"
	"chat-service/internal/http/utils/mappers"
	"chat-service/internal/utils/consts"
)

func (h *Handler) CreateChat(w http.ResponseWriter, r *http.Request) {
	var req request.CreateChat
	if err := helpers.DecodeJSON(r, &req); err != nil {
		helpers.HandleError(w, err)
		return
	}
	if req.Title == "" {
		helpers.HandleError(w, consts.InvalidChatTitle)
		return
	}

	ctx := helpers.WithTimeout(r.Context(), 500*time.Second)
	defer helpers.Cancel()

	chat := mappers.CreateChatToModel(req)
	if err := h.svc.CreateChat(ctx, &chat); err != nil {
		helpers.HandleError(w, err)
		return
	}

	chatResp := mappers.CreateChatToResponse(chat)
	helpers.SendJSON(w, http.StatusCreated, chatResp)
}

func (h *Handler) GetChatByID(w http.ResponseWriter, r *http.Request) {
	id := helpers.ParseParamID(w, r, "id")

	limit := helpers.QueryLimit(w, r)
	if limit == -1 {
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 500*time.Second)
	defer cancel()

	chat, err := h.svc.GetChatByID(ctx, id, limit)
	if err != nil {
		helpers.HandleError(w, err)
		return
	}

	chatResp := mappers.ChatToResponse(chat)
	helpers.SendJSON(w, http.StatusOK, chatResp)
}

func (h *Handler) DeleteChatByID(w http.ResponseWriter, r *http.Request) {
	id := helpers.ParseParamID(w, r, "id")

	ctx, cancel := context.WithTimeout(r.Context(), 500*time.Second)
	defer cancel()

	if err := h.svc.DeleteChatByID(ctx, id); err != nil {
		helpers.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

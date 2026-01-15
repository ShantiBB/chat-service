package handler

import (
	"context"
	"net/http"
	"strconv"
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

	ctx := helpers.WithTimeout(r.Context(), 500*time.Second)
	defer helpers.Cancel()

	chat := mappers.CreateChatToModel(req)
	if err := h.svc.CreateChat(ctx, &chat); err != nil {
		helpers.HandleError(w, err)
		return
	}

	chatResp := mappers.ChatToResponse(chat)
	helpers.SendJSON(w, http.StatusCreated, chatResp)
}

func (h *Handler) GetChatByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || id == 0 {
		helpers.HandleError(w, consts.InvalidChatID)
		return
	}

	limitParam := r.URL.Query().Get("limit")
	limit := 20
	if limitParam != "" {
		limit, err = strconv.Atoi(limitParam)
		if err != nil || limit < 0 {
			helpers.HandleError(w, consts.InvalidLimit)
			return
		}
	}

	ctx, cancel := context.WithTimeout(r.Context(), 500*time.Second)
	defer cancel()

	chat, err := h.svc.GetChatByID(ctx, uint(id), limit)
	if err != nil {
		helpers.HandleError(w, err)
		return
	}

	chatResp := mappers.ChatToResponse(chat)
	helpers.SendJSON(w, http.StatusOK, chatResp)
}

func (h *Handler) DeleteChatByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || id == 0 {
		helpers.HandleError(w, consts.InvalidChatID)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 500*time.Second)
	defer cancel()

	if err = h.svc.DeleteChatByID(ctx, uint(id)); err != nil {
		helpers.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

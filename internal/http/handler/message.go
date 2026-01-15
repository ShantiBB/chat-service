package handler

import (
	"net/http"
	"strconv"
	"time"

	"chat-service/internal/http/dto/request"
	"chat-service/internal/http/utils/helpers"
	"chat-service/internal/http/utils/mappers"
	"chat-service/internal/utils/consts"
)

func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var req request.CreateMessage
	if err := helpers.DecodeJSON(r, &req); err != nil {
		helpers.HandleError(w, err)
		return
	}

	chatIdParam := r.PathValue("chatID")
	chatID, err := strconv.ParseUint(chatIdParam, 10, 64)
	if err != nil || chatID == 0 {
		helpers.HandleError(w, consts.InvalidChatID)
		return
	}

	ctx := helpers.WithTimeout(r.Context(), 500*time.Second)
	defer helpers.Cancel()

	msg := mappers.CreateMessageToModel(req, uint(chatID))
	if err = h.svc.CreateMessage(ctx, &msg); err != nil {
		helpers.HandleError(w, err)
		return
	}

	msgResp := mappers.MessageToResponse(msg)
	helpers.SendJSON(w, http.StatusCreated, msgResp)
}

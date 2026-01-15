package handler

import (
	"net/http"
	"time"

	"chat-service/internal/http/dto/request"
	"chat-service/internal/http/utils/helpers"
	"chat-service/internal/http/utils/mappers"
	"chat-service/internal/http/utils/parsers"
	"chat-service/internal/utils/consts"
)

func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	chatID := parsers.ParseParamID(w, r, "chatID")
	req := request.CreateMessage{}
	if err := helpers.DecodeJSON(r, &req); err != nil {
		helpers.HandleError(w, err)
		return
	}
	if req.Text == "" {
		helpers.HandleError(w, consts.InvalidChatText)
		return
	}

	ctx := helpers.WithTimeout(r.Context(), 500*time.Second)
	defer helpers.Cancel()

	msg := mappers.CreateMessageToModel(req, chatID)
	if err := h.svc.CreateMessage(ctx, &msg); err != nil {
		helpers.HandleError(w, err)
		return
	}

	msgResp := mappers.MessageToResponse(msg)
	helpers.SendJSON(w, http.StatusCreated, msgResp)
}

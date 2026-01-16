package handler

import (
	"context"
	"net/http"

	"chat-service/internal/http/dto/request"
	"chat-service/internal/http/utils/helpers"
	"chat-service/internal/http/utils/mappers"
	"chat-service/internal/http/utils/parsers"
)

func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	chatID := parsers.ParseParamID(w, r, "chatID")
	req := request.CreateMessage{}
	if err := helpers.DecodeJSON(r, &req); err != nil {
		helpers.HandleError(w, err)
		return
	}

	err := req.Validate()
	if err != nil {
		helpers.HandleError(w, err)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), h.cfg.Server.Context.TimeOut)
	defer cancel()

	msg := mappers.CreateMessageToModel(&req, chatID)
	if err = h.svc.CreateMessage(ctx, msg); err != nil {
		helpers.HandleError(w, err)
		return
	}

	msgResp := mappers.MessageToResponse(msg)
	helpers.SendJSON(w, http.StatusCreated, msgResp)
}

package handler

import (
	"context"
	"net/http"

	"chat-service/internal/http/dto/request"
	_ "chat-service/internal/http/dto/response"
	"chat-service/internal/http/utils/helpers"
	"chat-service/internal/http/utils/mappers"
	"chat-service/internal/http/utils/parsers"
)

// CreateMessage   godoc
// @Summary      Create  a message
// @Description  Create a new message in chat
// @Tags         chats
// @Accept       json
// @Produce      json
// @Param		 chat_id	       path		   uint	                 true  "Chat id"
// @Param        request           body        request.CreateMessage true  "Message data"
// @Success      201               {object}    response.Message
// @Failure      400               {object}    helpers.apiError
// @Failure      404               {object}    helpers.apiError
// @Failure      500               {object}    helpers.apiError
// @Security     Bearer
// @Router       /chats/{chat_id}/messages            [post]
func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	chatID, ok := parsers.ParamID(w, r, "chatID")
	if !ok {
		return
	}

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

	ctx, cancel := context.WithTimeout(r.Context(), h.Cfg.Server.Context.TimeOut)
	defer cancel()

	msg := mappers.CreateMessageToModel(&req, chatID)
	if err = h.Svc.CreateMessage(ctx, msg); err != nil {
		helpers.HandleError(w, err)
		return
	}

	msgResp := mappers.MessageToResponse(msg)
	helpers.SendJSON(w, http.StatusCreated, msgResp)
}

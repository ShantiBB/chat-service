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

// CreateChat   godoc
// @Summary      Create a chat
// @Description  Create a new chat
// @Tags         chats
// @Accept       json
// @Produce      json
// @Param        request           body        request.CreateChat  true  "Chat data"
// @Success      201               {object}    response.Chat
// @Failure      400               {object}    helpers.apiError
// @Failure      500               {object}    helpers.apiError
// @Security     Bearer
// @Router       /chats            [post]
func (h *Handler) CreateChat(w http.ResponseWriter, r *http.Request) {
	var req request.CreateChat
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

	chat := mappers.CreateChatToModel(&req)
	if err = h.Svc.CreateChat(ctx, chat); err != nil {
		helpers.HandleError(w, err)
		return
	}

	chatResp := mappers.CreateChatToResponse(chat)
	helpers.SendJSON(w, http.StatusCreated, chatResp)
}

// GetChatByID    godoc
//
//	@Summary		Get a chat
//	@Description	Get a chat by id
//	@Tags			chats
//	@Accept			json
//	@Produce		json
//	@Param			chat_id	         path		     uint	true	"Chat id"
//	@Param			limit	         query		     int	false	"Limit messages" default(20)
//	@Success		200	{object}	 response.Chat
//	@Failure		400	{object}	 helpers.apiError
//	@Failure		404	{object}	 helpers.apiError
//	@Failure		500	{object}	 helpers.apiError
//	@Router			/chats/{chat_id} [get]
func (h *Handler) GetChatByID(w http.ResponseWriter, r *http.Request) {
	id, ok := parsers.ParamID(w, r, "id")
	if !ok {
		return
	}

	limit, ok := parsers.QueryLimit(w, r)
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), h.Cfg.Server.Context.TimeOut)
	defer cancel()

	chat, err := h.Svc.GetChatByID(ctx, id, limit)
	if err != nil {
		helpers.HandleError(w, err)
		return
	}

	chatResp := mappers.ChatToResponse(chat)
	helpers.SendJSON(w, http.StatusOK, chatResp)
}

// DeleteChatByID    godoc
//
//	@Summary		Delete a chat
//	@Description	Delete a chat by id
//	@Tags			chats
//	@Accept			json
//	@Produce		json
//	@Param			chat_id          path		     uint	true	"Chat id"
//	@Success		204	{object}	 nil
//	@Failure		400	{object}	 helpers.apiError
//	@Failure		404	{object}	 helpers.apiError
//	@Failure		500	{object}	 helpers.apiError
//	@Router			/chats/{chat_id} [delete]
func (h *Handler) DeleteChatByID(w http.ResponseWriter, r *http.Request) {
	id, ok := parsers.ParamID(w, r, "id")
	if !ok {
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), h.Cfg.Server.Context.TimeOut)
	defer cancel()

	if err := h.Svc.DeleteChatByID(ctx, id); err != nil {
		helpers.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

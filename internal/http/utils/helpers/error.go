package helpers

import (
	"errors"
	"net/http"

	"chat-service/internal/utils/consts"
)

//easyjson:json
type apiError struct {
	Error string `json:"error"`
}

func sendError(w http.ResponseWriter, code int, msg string) {
	SendJSON(w, code, apiError{Error: msg})
}

func HandleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, consts.InvalidChatID):
		sendError(w, http.StatusBadRequest, consts.MsgInvalidChatID)

	case errors.Is(err, consts.InvalidChatTitle):
		sendError(w, http.StatusBadRequest, consts.MsgInvalidChatTitle)

	case errors.Is(err, consts.InvalidMessagesLimit):
		sendError(w, http.StatusBadRequest, consts.MsgInvalidMessagesLimit)

	case errors.Is(err, consts.InvalidChatText):
		sendError(w, http.StatusBadRequest, consts.MsgInvalidChatText)

	case errors.Is(err, consts.ChatNotFound):
		sendError(w, http.StatusNotFound, consts.MsgChatNotFound)

	case errors.Is(err, consts.JsonEmptyBody):
		sendError(w, http.StatusBadRequest, consts.MsgJsonEmptyBody)

	case errors.Is(err, consts.JsonInvalid):
		sendError(w, http.StatusBadRequest, consts.MsgJsonInvalid)

	default:
		sendError(w, http.StatusInternalServerError, consts.MsgInternal)
	}
}

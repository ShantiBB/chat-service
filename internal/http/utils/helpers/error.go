package helpers

import (
	"errors"
	"net/http"

	ozzovalidation "github.com/go-ozzo/ozzo-validation"

	"chat-service/internal/http/utils/validation"
	"chat-service/internal/lib/utils/consts"
)

//easyjson:json
type apiError struct {
	Error string `json:"error"`
}

func SendError(w http.ResponseWriter, code int, msg string) {
	SendJSON(w, code, apiError{Error: msg})
}

func HandleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, consts.ErrInvalidChatID):
		SendError(w, http.StatusBadRequest, consts.MsgInvalidChatID)

	case errors.Is(err, consts.ErrInvalidMessagesLimit):
		SendError(w, http.StatusBadRequest, consts.MsgInvalidMessagesLimit)

	case errors.Is(err, consts.ErrChatNotFound):
		SendError(w, http.StatusNotFound, consts.MsgChatNotFound)

	case errors.Is(err, consts.ErrJsonEmptyBody):
		SendError(w, http.StatusBadRequest, consts.MsgJsonEmptyBody)

	case errors.Is(err, consts.ErrJsonInvalid):
		SendError(w, http.StatusBadRequest, consts.MsgJsonInvalid)

	case errors.As(err, &ozzovalidation.Errors{}):
		SendJSON(w, http.StatusBadRequest, validation.Error(err))

	default:
		SendError(w, http.StatusInternalServerError, consts.MsgInternal)
	}
}

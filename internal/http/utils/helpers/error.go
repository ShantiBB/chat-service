package helpers

import (
	"errors"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"

	"chat-service/internal/utils/consts"
)

//easyjson:json
type apiError struct {
	Error string `json:"error"`
}

//easyjson:json
type validateErrorResponse struct {
	Error  string            `json:"error"`
	Fields map[string]string `json:"fields"`
}

func SendError(w http.ResponseWriter, code int, msg string) {
	SendJSON(w, code, apiError{Error: msg})
}

func sendValidationError(w http.ResponseWriter, err error) {
	var validationErrs validation.Errors
	errors.As(err, &validationErrs)

	fieldErrors := make(map[string]string)
	for field, fieldErr := range validationErrs {
		fieldErrors[field] = fieldErr.Error()
	}

	response := validateErrorResponse{
		Error:  "validation failed",
		Fields: fieldErrors,
	}

	SendJSON(w, http.StatusBadRequest, response)
}

func HandleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, consts.InvalidChatID):
		SendError(w, http.StatusBadRequest, consts.MsgInvalidChatID)

	case errors.Is(err, consts.InvalidChatTitle):
		SendError(w, http.StatusBadRequest, consts.MsgInvalidChatTitle)

	case errors.Is(err, consts.InvalidMessagesLimit):
		SendError(w, http.StatusBadRequest, consts.MsgInvalidMessagesLimit)

	case errors.Is(err, consts.InvalidChatText):
		SendError(w, http.StatusBadRequest, consts.MsgInvalidChatText)

	case errors.Is(err, consts.ChatNotFound):
		SendError(w, http.StatusNotFound, consts.MsgChatNotFound)

	case errors.Is(err, consts.JsonEmptyBody):
		SendError(w, http.StatusBadRequest, consts.MsgJsonEmptyBody)

	case errors.Is(err, consts.JsonInvalid):
		SendError(w, http.StatusBadRequest, consts.MsgJsonInvalid)

	case errors.As(err, &validation.Errors{}):
		sendValidationError(w, err)

	default:
		SendError(w, http.StatusInternalServerError, consts.MsgInternal)
	}
}

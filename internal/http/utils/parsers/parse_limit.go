package parsers

import (
	"net/http"
	"strconv"

	"chat-service/internal/http/utils/helpers"
	"chat-service/internal/utils/consts"
)

const (
	DefaultLimit = 20
	MaxLimit     = 100
)

func QueryLimit(w http.ResponseWriter, r *http.Request) int {
	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		return DefaultLimit
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 0 || limit > MaxLimit {
		helpers.HandleError(w, consts.InvalidMessagesLimit)
		return -1
	}

	return limit
}

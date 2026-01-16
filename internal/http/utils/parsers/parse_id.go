package parsers

import (
	"net/http"
	"strconv"

	"chat-service/internal/http/utils/helpers"
	"chat-service/internal/lib/utils/consts"
)

func ParseParamID(w http.ResponseWriter, r *http.Request, param string) uint {
	idParam := r.PathValue(param)
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || id == 0 {
		helpers.HandleError(w, consts.ErrInvalidChatID)
		return 0
	}

	return uint(id)
}

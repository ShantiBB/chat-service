package helpers

import (
	"net/http"
	"strconv"

	"chat-service/internal/utils/consts"
)

func ParseParamID(w http.ResponseWriter, r *http.Request, param string) uint {
	idParam := r.PathValue(param)
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || id == 0 {
		HandleError(w, consts.InvalidChatID)
		return 0
	}

	return uint(id)
}

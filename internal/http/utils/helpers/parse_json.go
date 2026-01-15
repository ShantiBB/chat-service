package helpers

import (
	"io"
	"net/http"

	"github.com/mailru/easyjson"

	"chat-service/internal/utils/consts"
)

func SendJSON(w http.ResponseWriter, code int, v easyjson.Marshaler) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	b, err := easyjson.Marshal(v)
	if err != nil {
		return
	}
	_, _ = w.Write(b)
}

func DecodeJSON(r *http.Request, v easyjson.Unmarshaler) error {
	if r.Body == nil {
		return consts.JsonEmptyBody
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return consts.JsonInvalid
	}

	return easyjson.Unmarshal(b, v)
}

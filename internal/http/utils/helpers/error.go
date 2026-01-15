package helpers

import "net/http"

//easyjson:json
type apiError struct {
	Error string `json:"error"`
}

func SendError(w http.ResponseWriter, code int, msg string) {
	SendJSON(w, code, apiError{Error: msg})
}

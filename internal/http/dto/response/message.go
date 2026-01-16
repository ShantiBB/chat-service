package response

import "time"

//easyjson:json
type Message struct {
	CreatedAt time.Time `json:"created_at"`
	Text      string    `json:"text"`
	ID        uint      `json:"id"`
}

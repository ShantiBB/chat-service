package response

import "time"

//easyjson:json
type Message struct {
	ID        uint      `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

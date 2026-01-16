package response

import "time"

//easyjson:json
type CreateChat struct {
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	ID        uint      `json:"id"`
}

//easyjson:json
type Chat struct {
	CreatedAt time.Time  `json:"created_at"`
	Title     string     `json:"title"`
	Messages  []*Message `json:"messages"`
	ID        uint       `json:"id"`
}

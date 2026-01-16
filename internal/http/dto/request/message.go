package request

import (
	validation "github.com/go-ozzo/ozzo-validation"

	"chat-service/internal/lib/utils/consts"
)

//easyjson:json
type CreateMessage struct {
	Text string `json:"text"`
}

func (c *CreateMessage) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(
			&c.Text,
			validation.Required.Error(consts.MsgRequiredField),
			validation.Length(1, 5000).Error(consts.MsgTextLength),
		),
	)
}

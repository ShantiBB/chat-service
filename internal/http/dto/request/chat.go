package request

import (
	validation "github.com/go-ozzo/ozzo-validation"

	"chat-service/internal/lib/utils/consts"
)

//easyjson:json
type CreateChat struct {
	Title string `json:"title"`
}

func (c *CreateChat) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(
			&c.Title,
			validation.Required.Error(consts.MsgRequiredField),
			validation.Length(1, 5000).Error(consts.MsgTitleLength),
		),
	)
}

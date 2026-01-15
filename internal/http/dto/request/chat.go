package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
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
			validation.Required.Error("название чата обязательно"),
			validation.Length(1, 200).Error("название должно быть от 1 до 200 символов"),
		),
	)
}

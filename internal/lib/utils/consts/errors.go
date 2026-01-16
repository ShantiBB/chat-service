package consts

import "errors"

const (
	MsgInvalidChatID        = "invalid chat id"
	MsgRequiredField        = "field is required"
	MsgInvalidMessagesLimit = "limit query must be between 1 and 100"
	MsgTextLength           = "text length must be between 1 and 5000"
	MsgTitleLength          = "title length must be between 1 and 200"
	MsgChatNotFound         = "chat not found"
	MsgJsonEmptyBody        = "empty json body"
	MsgJsonInvalid          = "invalid json"
	MsgInternal             = "internal server error"
)

var (
	ErrInvalidChatID        = errors.New(MsgInvalidChatID)
	ErrInvalidMessagesLimit = errors.New(MsgInvalidMessagesLimit)
	ErrChatNotFound         = errors.New(MsgChatNotFound)
	ErrJsonEmptyBody        = errors.New(MsgJsonEmptyBody)
	ErrJsonInvalid          = errors.New(MsgJsonInvalid)
)

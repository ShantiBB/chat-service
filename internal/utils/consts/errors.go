package consts

import "errors"

const (
	MsgInvalidChatID        = "invalid chat id"
	MsgInvalidChatTitle     = "invalid chat title (must be not empty)"
	MsgInvalidMessagesLimit = "invalid limit query"
	MsgInvalidChatText      = "invalid message text (must be not empty)"
	MsgChatNotFound         = "chat not found"
	MsgJsonEmptyBody        = "empty json body"
	MsgJsonInvalid          = "invalid json"
	MsgInternal             = "internal server error"
)

var (
	InvalidChatID        = errors.New(MsgInvalidChatID)
	InvalidChatTitle     = errors.New(MsgInvalidMessagesLimit)
	InvalidMessagesLimit = errors.New(MsgInvalidMessagesLimit)
	InvalidChatText      = errors.New(MsgInvalidMessagesLimit)
	ChatNotFound         = errors.New(MsgChatNotFound)
	JsonEmptyBody        = errors.New(MsgJsonEmptyBody)
	JsonInvalid          = errors.New(MsgJsonInvalid)
)

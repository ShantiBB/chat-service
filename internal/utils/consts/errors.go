package consts

import "errors"

const (
	MsgInvalidChatID = "invalid chat id"
	MsgInvalidLimit  = "invalid limit query"
	MsgChatNotFound  = "chat not found"
	MsgJsonEmptyBody = "empty json body"
	MsgJsonInvalid   = "invalid json"
	MsgInternal      = "internal server error"
)

var (
	InvalidChatID = errors.New(MsgInvalidChatID)
	InvalidLimit  = errors.New(MsgInvalidLimit)
	ChatNotFound  = errors.New(MsgChatNotFound)
	JsonEmptyBody = errors.New(MsgJsonEmptyBody)
	JsonInvalid   = errors.New(MsgJsonInvalid)
)

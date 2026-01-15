package consts

import "errors"

var (
	ChatNotFound     = errors.New("chat not found")
	FailedDeleteChat = errors.New("failed delete chat")
)

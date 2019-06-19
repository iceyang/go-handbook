package internal

import "errors"

type Error struct {
	Err  error
	Msg  string
	Code int
}

func NewError(code int) {
	panic(&Error{
		Err:  errors.New("Request Error"),
		Msg:  "Request Error",
		Code: code,
	})
}

func NewErrorWithMsg(code int, msg string) {
	panic(&Error{
		Err:  errors.New(msg),
		Msg:  msg,
		Code: code,
	})
}

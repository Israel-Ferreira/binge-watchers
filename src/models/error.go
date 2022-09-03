package models

type ErrorResp struct {
	Msg string `json:"msg"`
}

func NewErrorResp(msg string) ErrorResp {
	return ErrorResp{Msg: msg}
}

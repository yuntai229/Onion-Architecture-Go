package model

import "onion-architecrure-go/domain/constant"

type ResponseModel struct{}

type ResSucc struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResFail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewResModel() ResponseModel {
	return ResponseModel{}
}

func (model ResponseModel) ResWithSucc(data any) ResSucc {
	res := ResSucc{
		Code:    "0000",
		Message: "Succ",
		Data:    data,
	}
	return res
}

func (model ResponseModel) ResWithFail(err constant.ErrorMessage) ResFail {
	res := ResFail{
		Code:    err.Code,
		Message: err.Message,
	}
	return res
}

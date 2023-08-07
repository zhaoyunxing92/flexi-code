package handler

import (
	"github.com/zhaoyunxing92/flexi-code/server/internal/translator"
)

type Body interface{}

type RespBody[T Body] struct {
	// http code
	Code int `json:"code"`

	// 是否成功
	Succeed bool `json:"succeed"`

	// response message
	Message string `json:"msg"`

	// reason key
	Reason string `json:"reason"`

	// response data
	Data Body `json:"data,omitempty"`
}

func (res *RespBody[T]) TrMsg(lang translator.Language) {
	if len(res.Message) == 0 {
		res.Message = translator.Tr(lang, res.Reason)
	}
}

func New[T Body](code int, succeed bool, msg, reason string, data T) RespBody[T] {
	return RespBody[T]{code, succeed, msg, reason, data}
}

func Succeed(msg, reason string, data Body) RespBody[Body] {
	return New[Body](200, true, msg, reason, data)
}

func Fail(code int, msg, reason string) RespBody[Body] {
	return New[Body](code, false, msg, reason, nil)
}

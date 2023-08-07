package handler

type Body interface{}

type Resp[T Body] struct {
	// http code
	Code int `json:"code"`

	// 是否成功
	Succeed bool `json:"succeed"`

	// response message
	Message string `json:"msg"`

	// response data
	Data Body `json:"data,omitempty"`
}

// TrMsg translate the reason cause as a message
//func (r *Resp) TrMsg(lang i18n.Language) *Resp {
//	if len(r.Message) == 0 {
//		r.Message = translator.Tr(lang, r.Reason)
//	}
//	return r
//}

func New[T Body](code int, succeed bool, msg string, data T) Resp[T] {
	return Resp[T]{code, succeed, msg, data}
}

func Succeed(msg string, data Body) Resp[Body] {
	return New[Body](200, true, msg, data)
}

func Fail(code int, msg, reason string) Resp[Body] {
	return New[Body](code, false, msg, reason)
}

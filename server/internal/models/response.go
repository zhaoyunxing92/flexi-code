package models

type Body interface{}

type Response[T Body] struct {
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
func (r *Response) TrMsg(lang i18n.Language) *RespBody {
	if len(r.Message) == 0 {
		r.Message = translator.Tr(lang, r.Reason)
	}
	return r
}

func New[T Body](code int, succeed bool, msg string, data T) Response[T] {
	return Response[T]{code, succeed, msg, data}
}

func Succeed(msg string, data Body) Response[Body] {
	return New[Body](200, true, msg, data)
}

func Fail(code int, msg, reason string) Response[Body] {
	return New[Body](code, false, msg, reason)
}

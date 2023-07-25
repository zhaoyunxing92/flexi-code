package errors

type Error struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Reason string `json:"reason"`
}

// NewError create error
func NewError(code int, msg, reason string) Error {
	return Error{Code: code, Msg: msg, Reason: reason}
}

// New create error
func New(msg string) Error {
	return Error{Code: 500, Msg: msg, Reason: "sys"}
}

// Error return error with info
func (e *Error) Error() string {
	return e.Msg
}

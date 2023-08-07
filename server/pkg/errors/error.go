package errors

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

type Error struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Reason string `json:"reason"`
	Err    error
	Stack  string
}

// NewError create error
func NewError(code int, msg, reason string) *Error {
	return &Error{Code: code, Msg: msg, Reason: reason}
}

// New create error
func New(msg string) *Error {
	return &Error{Code: 500, Msg: msg, Reason: "sys"}
}

// Error return error with info
func (e *Error) Error() string {
	return e.Msg
}

// WithError with original error
func (e *Error) WithError(err error) *Error {
	e.Err = err
	return e
}

// WithStack with stack
func (e *Error) WithStack() *Error {
	e.Stack = logStack(2, 0)
	return e
}

func (e *Error) Format(state fmt.State, verb rune) {
	switch verb {
	case 'v':
		str := bytes.NewBuffer([]byte{})
		str.WriteString(fmt.Sprintf("code: %d, ", e.Code))
		str.WriteString("reason: ")
		str.WriteString(e.Reason + ", ")
		str.WriteString("message: ")
		str.WriteString(e.Msg)
		if e.Err != nil {
			str.WriteString(", error: ")
			str.WriteString(e.Err.Error())
		}
		if len(e.Stack) > 0 {
			str.WriteString("\n")
			str.WriteString(e.Stack)
		}
		fmt.Fprintf(state, "%s", strings.Trim(str.String(), "\r\n\t"))
	default:
		fmt.Fprintf(state, e.Msg)
	}
}

func logStack(start, end int) string {
	stack := bytes.Buffer{}
	for i := start; i < end || end <= 0; i++ {
		pc, str, line, _ := runtime.Caller(i)
		if line == 0 {
			break
		}
		stack.WriteString(fmt.Sprintf("%s:%d %s\n", str, line, runtime.FuncForPC(pc).Name()))
	}
	return stack.String()
}

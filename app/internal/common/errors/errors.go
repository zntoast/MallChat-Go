package errors

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

// 定义错误码
const (
	Success         = 0
	UnknownError    = 10001
	ValidationError = 10002
	AuthError       = 10003
	NotFoundError   = 10004
)

// Error 自定义错误结构
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("{ErrCode:%d,ErrMsg:%s}", e.Code, e.Message)
}

// New 创建新的错误
func New(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func Warp(code int, err error, message string) *Error {
	e := New(code, message)
	logx.ErrorStack(e, e.Error())
	return e
}

// NewValidationError 创建验证错误
func NewValidationError(message string) *Error {
	return New(ValidationError, message)
}

// NewAuthError 创建认证错误
func NewAuthError(message string) *Error {
	return New(AuthError, message)
}

// NewNotFoundError 创建未找到错误
func NewNotFoundError(message string) *Error {
	return New(NotFoundError, message)
}

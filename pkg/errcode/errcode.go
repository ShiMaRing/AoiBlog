package errcode

import (
	"fmt"
	"log"
	"net/http"
)

// Error omitempty表示不存在的时候直接忽略，但是不能够忽略结构体
type Error struct {
	code    int      `json:"code,omitempty"`
	msg     string   `json:"msg,omitempty"`
	details []string `json:"details,omitempty"`
}

//使用自定义码转http码，避免重复
var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		log.Fatalln("错误码已存在，请重新定义")
	}
	codes[code] = msg
	return &Error{
		code: code,
		msg:  msg,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d ,错误信息： %s \n", e.Code(), e.Msg())
}

func (e *Error) StatusCode() int {
	switch e.code {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Code() int {
	return e.code
}
func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetail(details ...string) *Error {
	newErr := *e //不应该使用指针，应该是有结构体，避免修改原始数据
	newErr.details = []string{}
	for i := range details {
		newErr.details = append(newErr.details, details[i])
	}
	return &newErr
}
func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

package status

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gen/api/codes"
	"net/http"
)

type Status struct {
	ctx   *gin.Context
	C     codes.Code    `json:"code"`
	M     string        `json:"message"`
	E     string        `json:"error"` // 错误提示信息
	mArgs []interface{} // 如果错误码对应的消息是一个fmt模板, 则可以使用该参数格式化消息
}

type Debug struct {
	Hostname string `json:"hostname"`
	Time     string `json:"time"`
}

func New(code codes.Code) *Status {
	return &Status{
		C: code,
	}
}

// Newf 如果错误码对应的消息格式是一个fmt模板, 则可以使用该参数格式化消息
func Newf(code codes.Code, a ...interface{}) *Status {
	return &Status{
		C:     code,
		mArgs: a,
	}
}

func (s *Status) withCtx(c *gin.Context) *Status {
	s.ctx = c
	return s
}

func (s *Status) WithErrorf(format string, a ...interface{}) *Status {
	s.E = fmt.Sprintf(format, a...)
	return s
}

func (s *Status) WithError(a ...interface{}) *Status {
	s.E = fmt.Sprint(a...)
	return s
}

func (s *Status) Error() string {
	return fmt.Sprintf("httpStatus = %d; code = %v; msg = %s; err = %s", s.httpCode(), s.C, s.M, s.E)
}

func (s *Status) Code() codes.Code {
	if s == nil {
		return codes.OK
	}
	return s.C
}

/*
// 根据用户自定义的错误码 映射HTTP code
// 自定义错误码编码格式
//
// 错误类型(1位) & 模块(2位) & 错误编码(3位)
//   4            1 0        1 0 6
//
// 映射规则
// 1. 不篡改http status语义, 当自定义code是一个标准的http code时则原样返回, 返回的结构体中code与http status保持一致
// 2. 取自定义code中的第1位, 进行http status映射
//    - 2 表示不是错误   映射 http.StatusOK  前端可以根据code做判断或者跳转用
//    - 4 表示用户错误   映射 http.StatusBadRequest
//    - 5 表示服务错误   映射 http.StatusInternalServerError
*/
func (s *Status) httpCode() int {
	code := s.Code()
	switch code.GetErrorType() {
	case "5":
		return http.StatusInternalServerError
	case "4":
		return http.StatusBadRequest
	case "2":
		return http.StatusOK
	default:
		return http.StatusBadRequest
	}
}

func (s *Status) Message() string {
	if s == nil {
		return ""
	}
	return s.M
}

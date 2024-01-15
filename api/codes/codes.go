package codes

import (
	"net/http"
	"strconv"
)

type Code int64

// ErrorCodePrefix  服务错误码前2位号段代表模塊名稱
type ErrorCodePrefix int64

const (
	OK           Code = 200
	BadRequest   Code = 400
	Unauthorized Code = 401
	Forbidden    Code = 403
	NotFound     Code = 404
	Internal     Code = 500
)

var CodeToStr = map[Code]string{
	OK:           "成功",
	BadRequest:   "请求错误",
	Unauthorized: "未授权的请求", // 用户鉴权不通过 token校验失败
	Forbidden:    "没有权限访问", // 用户鉴权通过了 但是对资源没有权限访问
	NotFound:     "请求资源不存在",
	Internal:     "内部服务器错误",
}

func Message(code Code) string {
	if s, ok := CodeToStr[code]; ok {
		return s
	}
	return "该错误码未定义错误信息"
}

// Int int轉換
func (c Code) Int() int {
	return int(c)
}

// ToString 字符串轉換
func (c Code) ToString() string {
	return strconv.FormatInt(int64(c), 10)
}

func (c Code) Message() string {
	return Message(c)
}

// StandardHTTPStatus 判定是否是标准的HTTP code
func (c Code) StandardHTTPStatus() bool {
	// code是一个标准的http code
	if http.StatusText(c.Int()) != "" {
		return true
	}
	return false
}

// Customize 自定义错误码
func (c Code) Customize() bool {
	if len(c.ToString()) == 6 {
		return true
	}
	return false
}

// GetErrorType 获取错误码类型 第5位
func (c Code) GetErrorType() string {
	return c.ToString()[0:1]
}

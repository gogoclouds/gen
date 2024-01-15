package body

import (
	"github.com/gogoclouds/gen/api/codes"
	"time"
)

type PageData[T any] struct {
	Total int64 `json:"total"`
	List  []T   `json:"list"`
}

func NewPageData[T any](total int64) *PageData[T] {
	return &PageData[T]{
		Total: total,
		List:  make([]T, 0, total),
	}
}

type Response[T any] struct {
	Code int    `json:"code"`
	Data T      `json:"data"`
	Msg  string `json:"msg"`
	Time int64  `json:"time"`
}

// NewResponse 响应成功
func NewResponse[T any](data T) *Response[T] {
	return &Response[T]{
		Code: 0,
		Msg:  "",
		Data: data,
		Time: time.Now().Unix(),
	}
}

func NewResponseErr[T any](code codes.Code, data T) *Response[T] {
	return &Response[T]{
		Code: code.Int(),
		Msg:  code.Message(),
		Data: data,
		Time: time.Now().Unix(),
	}
}

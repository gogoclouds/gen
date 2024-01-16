package v1

import (
	"github.com/gogoclouds/gen/api/body"
	"github.com/gogoclouds/gen/internal/model"
)

type UserListRequest struct {
	body.PageRequest
}

type UserListResponse struct {
	body.PageData[*model.SysUser]
}

type UserRequest struct {
	ID uint32 `json:"id"`
}

type UserResponse struct {
	*model.SysUser
}

type UserCreateRequest struct {
}

type UserCreateResponse struct {
}

type UserUpdateRequest struct {
	ID uint32 `json:"id"`
}

type UserUpdateResponse struct {
}

type UserDeleteRequest struct {
	ID uint32 `json:"id"`
}

type UserDeleteResponse struct {
}

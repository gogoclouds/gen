package v1

import (
	"github.com/gogoclouds/gen/api/body"
	"github.com/gogoclouds/gen/internal/model"
)

type SampleListRequest struct {
	body.PageRequest
}

type SampleListResponse struct {
	body.PageData[*model.SysUser]
}

type SampleRequest struct {
	ID uint32 `json:"id"`
}

type SampleResponse struct {
	*model.SysUser
}

type SampleCreateRequest struct {
}

type SampleCreateResponse struct {
}

type SampleUpdateRequest struct {
	ID uint32 `json:"id"`
}

type SampleUpdateResponse struct {
}

type SampleDeleteRequest struct {
	ID uint32 `json:"id"`
}

type SampleDeleteResponse struct {
}

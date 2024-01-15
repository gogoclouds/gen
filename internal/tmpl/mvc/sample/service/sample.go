package service

import (
	"context"
	"github.com/gogoclouds/gen/api/body"
	"github.com/gogoclouds/gen/internal/model"
	"github.com/gogoclouds/gen/internal/tmpl/mvc/sample/api/v1"
	"github.com/jinzhu/copier"
)

type ISampleRepo interface {
	Find(ctx context.Context, req *v1.SampleListRequest) ([]*model.SysUser, int64, error)
	FindOne(ctx context.Context, req *v1.SampleRequest) (*model.SysUser, error)
	Create(ctx context.Context, data *model.SysUser) error
	Update(ctx context.Context, req *v1.SampleUpdateRequest) error
	Delete(ctx context.Context, req *v1.SampleDeleteRequest) error
}

type SampleService struct {
	repo ISampleRepo
}

func NewSampleService(repo ISampleRepo) *SampleService {
	return &SampleService{repo: repo}
}

func (svc *SampleService) List(ctx context.Context, req *v1.SampleListRequest) (*body.PageData[*model.SysUser], error) {
	list, total, err := svc.repo.Find(ctx, req)
	if err != nil {
		return nil, err
	}
	return &body.PageData[*model.SysUser]{
		Total: total,
		List:  list,
	}, nil
}

func (svc *SampleService) GetDetails(ctx context.Context, req *v1.SampleRequest) (*v1.SampleResponse, error) {
	one, err := svc.repo.FindOne(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.SampleResponse{SysUser: one}, nil
}

func (svc *SampleService) Add(ctx context.Context, req *v1.SampleCreateRequest) error {
	var data model.SysUser
	copier.Copy(&data, req)
	return svc.repo.Create(ctx, &data)
}

func (svc *SampleService) Update(ctx context.Context, req *v1.SampleUpdateRequest) error {
	return svc.repo.Update(ctx, req)
}

func (svc *SampleService) Delete(ctx context.Context, req *v1.SampleDeleteRequest) error {
	return svc.repo.Delete(ctx, req)
}

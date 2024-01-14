package service

import (
	"context"
	"github.com/gogoclouds/gen/internal/tmpl/mvc/sample/api/v1"
)

type ISampleRepo interface {
	Find(ctx context.Context, req *v1.SampleListRequest) (*v1.SampleListResponse, error)
	FindOne(ctx context.Context, req *v1.SampleRequest) (*v1.SampleResponse, error)
	Create(ctx context.Context, req *v1.SampleCreateRequest) (*v1.SampleCreateResponse, error)
	Update(ctx context.Context, req *v1.SampleUpdateRequest) (*v1.SampleUpdateResponse, error)
	Delete(ctx context.Context, req *v1.SampleDeleteRequest) (*v1.SampleDeleteResponse, error)
}

type SampleService struct {
	repo ISampleRepo
}

func NewSampleService(repo ISampleRepo) *SampleService {
	return &SampleService{repo: repo}
}

func (svc *SampleService) List(ctx context.Context, req *v1.SampleListRequest) (*v1.SampleListResponse, error) {
	return svc.repo.Find(ctx, req)
}

func (svc *SampleService) GetDetails(ctx context.Context, req *v1.SampleRequest) (*v1.SampleResponse, error) {
	return svc.repo.FindOne(ctx, req)
}

func (svc *SampleService) Add(ctx context.Context, req *v1.SampleCreateRequest) (*v1.SampleCreateResponse, error) {
	return svc.repo.Create(ctx, req)
}

func (svc *SampleService) Update(ctx context.Context, req *v1.SampleUpdateRequest) (*v1.SampleUpdateResponse, error) {
	return svc.repo.Update(ctx, req)
}

func (svc *SampleService) Delete(ctx context.Context, req *v1.SampleDeleteRequest) (*v1.SampleDeleteResponse, error) {
	return svc.repo.Delete(ctx, req)
}

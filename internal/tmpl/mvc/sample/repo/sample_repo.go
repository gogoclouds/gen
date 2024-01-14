package repo

import (
	"context"
	"github.com/gogoclouds/gen/internal/tmpl/mvc/sample/api/v1"
	"gorm.io/gorm"
)

type SampleRepo struct {
	db *gorm.DB
	// q *qurey.Query
}

func NewSampleRepo(db *gorm.DB) *SampleRepo {
	return &SampleRepo{db: db}
}

func (repo *SampleRepo) Find(ctx context.Context, req *v1.SampleListRequest) (*v1.SampleListResponse, error) {
	return nil, nil
}

func (repo *SampleRepo) FindOne(ctx context.Context, req *v1.SampleRequest) (*v1.SampleResponse, error) {
	return nil, nil
}

func (repo *SampleRepo) Create(ctx context.Context, req *v1.SampleCreateRequest) (*v1.SampleCreateResponse, error) {
	return nil, nil
}

func (repo *SampleRepo) Update(ctx context.Context, req *v1.SampleUpdateRequest) (*v1.SampleUpdateResponse, error) {
	return nil, nil
}

func (repo *SampleRepo) Delete(ctx context.Context, req *v1.SampleDeleteRequest) (*v1.SampleDeleteResponse, error) {
	return nil, nil
}

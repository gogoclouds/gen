package repo

import (
	"context"
	"github.com/gogoclouds/gen/internal/model"
	"github.com/gogoclouds/gen/internal/query"
	"github.com/gogoclouds/gen/internal/tmpl/mvc/sample/api/v1"
	"gorm.io/gorm"
)

type SampleRepo struct {
	db *gorm.DB
	q  *query.Query
}

func NewSampleRepo(db *gorm.DB) *SampleRepo {
	return &SampleRepo{db: db, q: query.Use(db)}
}

func (repo *SampleRepo) Find(ctx context.Context, req *v1.SampleListRequest) (result []*model.SysUser, count int64, err error) {
	q := repo.q.SysUser
	return q.WithContext(ctx).FindByPage(req.ToOffset(), req.ToLimit())
}

func (repo *SampleRepo) FindOne(ctx context.Context, req *v1.SampleRequest) (*model.SysUser, error) {
	q := repo.q.SysUser
	return q.WithContext(ctx).Where(q.ID.Eq(req.ID)).First()
}

func (repo *SampleRepo) Create(ctx context.Context, data *model.SysUser) error {
	return repo.q.SysUser.WithContext(ctx).Create(data)
}

func (repo *SampleRepo) Update(ctx context.Context, req *v1.SampleUpdateRequest) error {
	q := repo.q.SysUser
	_, err := q.WithContext(ctx).Where(q.ID.Eq(req.ID)).Updates(req)
	return err
}

func (repo *SampleRepo) Delete(ctx context.Context, req *v1.SampleDeleteRequest) error {
	q := repo.q.SysUser
	_, err := q.WithContext(ctx).Where(q.ID.Eq(req.ID)).Delete()
	return err
}

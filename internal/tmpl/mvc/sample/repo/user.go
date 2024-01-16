package repo

import (
	"context"
	"github.com/gogoclouds/gen/internal/model"
	"github.com/gogoclouds/gen/internal/query"
	"github.com/gogoclouds/gen/internal/tmpl/mvc/sample/api/v1"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
	q  *query.Query
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db, q: query.Use(db)}
}

func (repo *UserRepo) Find(ctx context.Context, req *v1.UserListRequest) (result []*model.SysUser, count int64, err error) {
	q := repo.q.SysUser
	return q.WithContext(ctx).FindByPage(req.ToOffset(), req.ToLimit())
}

func (repo *UserRepo) FindOne(ctx context.Context, req *v1.UserRequest) (*model.SysUser, error) {
	q := repo.q.SysUser
	return q.WithContext(ctx).Where(q.ID.Eq(req.ID)).First()
}

func (repo *UserRepo) Create(ctx context.Context, data *model.SysUser) error {
	return repo.q.SysUser.WithContext(ctx).Create(data)
}

func (repo *UserRepo) Update(ctx context.Context, req *v1.UserUpdateRequest) error {
	q := repo.q.SysUser
	_, err := q.WithContext(ctx).Where(q.ID.Eq(req.ID)).Updates(req)
	return err
}

func (repo *UserRepo) Delete(ctx context.Context, req *v1.UserDeleteRequest) error {
	q := repo.q.SysUser
	_, err := q.WithContext(ctx).Where(q.ID.Eq(req.ID)).Delete()
	return err
}

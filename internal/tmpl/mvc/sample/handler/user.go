package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gen/api"
	"github.com/gogoclouds/gen/api/codes"
	"github.com/gogoclouds/gen/api/status"
	"github.com/gogoclouds/gen/internal/tmpl/mvc/sample/api/v1"
	"github.com/gogoclouds/gen/internal/tmpl/mvc/sample/service"
	"log/slog"
)

type UserServer struct {
	svc *service.UserService
}

func NewUserServer(svc *service.UserService) *UserServer {
	return &UserServer{svc: svc}
}

func (h *UserServer) List(ctx *gin.Context) {
	req := new(v1.UserListRequest)
	if err := ctx.ShouldBind(req); err != nil {
		api.Fail(ctx, status.New(codes.BadRequest).WithError(err))
		return
	}
	list, err := h.svc.List(ctx, req)
	if err != nil {
		slog.ErrorContext(ctx, "List error", slog.String("err", err.Error()))
		api.Fail(ctx, status.New(codes.Internal))
		return
	}
	api.OK(ctx, list)
}

func (h *UserServer) Details(ctx *gin.Context) {
	req := new(v1.UserRequest)
	if err := ctx.ShouldBind(req); err != nil {
		api.Fail(ctx, status.New(codes.BadRequest).WithError(err))
		return
	}
	data, err := h.svc.GetDetails(ctx, req)
	if err != nil {
		slog.ErrorContext(ctx, "Details error", slog.String("err", err.Error()))
		api.Fail(ctx, status.New(codes.Internal))
		return
	}
	api.OK(ctx, data)
}

func (h *UserServer) Add(ctx *gin.Context) {
	req := new(v1.UserCreateRequest)
	if err := ctx.ShouldBind(req); err != nil {
		api.Fail(ctx, status.New(codes.BadRequest).WithError(err))
		return
	}
	if err := h.svc.Add(ctx, req); err != nil {
		slog.ErrorContext(ctx, "Add error", slog.String("err", err.Error()))
		api.Fail(ctx, status.New(codes.Internal))
		return
	}
	api.OK(ctx, gin.H{})
}

func (h *UserServer) Update(ctx *gin.Context) {
	req := new(v1.UserUpdateRequest)
	if err := ctx.ShouldBind(req); err != nil {
		api.Fail(ctx, status.New(codes.BadRequest).WithError(err))
		return
	}
	if err := h.svc.Update(ctx, req); err != nil {
		slog.ErrorContext(ctx, "Update error", slog.String("err", err.Error()))
		api.Fail(ctx, status.New(codes.Internal))
		return
	}
	api.OK(ctx, gin.H{})
}

func (h *UserServer) Delete(ctx *gin.Context) {
	req := new(v1.UserDeleteRequest)
	if err := ctx.ShouldBind(req); err != nil {
		api.Fail(ctx, status.New(codes.BadRequest).WithError(err))
		return
	}
	if err := h.svc.Delete(ctx, req); err != nil {
		slog.ErrorContext(ctx, "Delete error", slog.String("err", err.Error()))
		api.Fail(ctx, status.New(codes.Internal))
		return
	}
	api.OK(ctx, gin.H{})
}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gen/internal/tmpl/mvc/sample/api/v1"
	"github.com/gogoclouds/gen/internal/tmpl/mvc/sample/service"
	"net/http"
)

type SampleServer struct {
	svc *service.SampleService
}

func NewSampleServer(svc *service.SampleService) *SampleServer {
	return &SampleServer{svc: svc}
}

func (h *SampleServer) List(ctx *gin.Context) {
	req := new(v1.SampleListRequest)
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	list, err := h.svc.List(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  "success",
		"data": list,
	})
}

func (h *SampleServer) Details(ctx *gin.Context) {
	req := new(v1.SampleRequest)
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	data, err := h.svc.GetDetails(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  "success",
		"data": data,
	})
}

func (h *SampleServer) Add(ctx *gin.Context) {
	req := new(v1.SampleCreateRequest)
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	list, err := h.svc.Add(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  "success",
		"data": list,
	})
}

func (h *SampleServer) Update(ctx *gin.Context) {
	req := new(v1.SampleUpdateRequest)
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	list, err := h.svc.Update(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  "success",
		"data": list,
	})
}

func (h *SampleServer) Delete(ctx *gin.Context) {
	req := new(v1.SampleDeleteRequest)
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	list, err := h.svc.Delete(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  "success",
		"data": list,
	})
}

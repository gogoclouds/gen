package sample

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gen/internal/tmpl/mvc/sample/handler"
	"github.com/gogoclouds/gen/internal/tmpl/mvc/sample/repo"
	"github.com/gogoclouds/gen/internal/tmpl/mvc/sample/service"
)

func Register(r gin.IRouter) {
	sampleGp := r.Group("/sample")
	{
		sampleHandler := handler.NewSampleServer(service.NewSampleService(repo.NewSampleRepo(nil)))
		sampleGp.POST("/list", sampleHandler.List)
		sampleGp.POST("/details", sampleHandler.Details)
		sampleGp.POST("/add", sampleHandler.Add)
		sampleGp.POST("/update", sampleHandler.Update)
		sampleGp.POST("/delete", sampleHandler.Delete)
	}
}

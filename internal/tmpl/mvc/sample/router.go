package sample

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gen/internal/tmpl/mvc/sample/handler"
	"github.com/gogoclouds/gen/internal/tmpl/mvc/sample/repo"
	"github.com/gogoclouds/gen/internal/tmpl/mvc/sample/service"
)

func Register(r gin.IRouter) {
	userGp := r.Group("/user")
	{
		userHandler := handler.NewUserServer(service.NewUserService(repo.NewUserRepo(nil)))
		userGp.POST("/list", userHandler.List)
		userGp.POST("/details", userHandler.Details)
		userGp.POST("/add", userHandler.Add)
		userGp.POST("/update", userHandler.Update)
		userGp.POST("/delete", userHandler.Delete)
	}
}

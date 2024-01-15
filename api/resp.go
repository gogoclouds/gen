package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gogoclouds/gen/api/body"
	"github.com/gogoclouds/gen/api/status"
	"net/http"
)

func OK[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, body.NewResponse(data))
}

func Fail(c *gin.Context, status *status.Status) {
	c.JSON(http.StatusOK, body.NewResponseErr(status.Code(), status.E))
}

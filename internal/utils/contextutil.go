package utils

import (
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/hanakogo/exceptiongo"
	"net/http"
)

func CtxBindJson[T any](c *gin.Context) *T {
	target := NewValuePtr[T]()
	err := c.ShouldBindJSON(target)
	exceptiongo.ThrowErr[types.JsonBindingError](err)
	return target
}

func CtxBindQuery[T any](c *gin.Context) *T {
	target := NewValuePtr[T]()
	err := c.ShouldBindQuery(target)
	exceptiongo.ThrowErr[types.QueryBindingError](err)
	return target
}

type Response[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}

func CtxRespOK[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, Response[T]{http.StatusOK, data})
}

func CtxRespCustom[T any](c *gin.Context, code int, data T) {
	c.JSON(code, Response[T]{code, data})
}

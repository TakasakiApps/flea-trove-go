package middlewares

import (
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/hanakogo/exceptiongo"
	"github.com/hanakogo/hanakoutilgo"
)

func getErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if hanakoutilgo.Is[error](r) {
					exceptiongo.ThrowErr[types.HttpUnknownError](hanakoutilgo.CastTo[error](r))
				}
			}
		}()
	}
}

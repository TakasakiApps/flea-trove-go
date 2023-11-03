package middlewares

import (
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/TakasakiApps/flea-trove-go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/hanakogo/exceptiongo"
	"net/http"
)

func getHttpExceptionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 部署Http异常处理器
		defer exceptiongo.NewExceptionHandler(func(e *exceptiongo.Exception) {
			customResponse := func(code int) {
				utils.CtxRespCustom[string](c, code, e.Error().Error())
			}
			c.Abort()
			logger.Error(e.GetStackTraceMessage())
			switch e.Type() {
			case exceptiongo.TypeOf[types.StatusConflict]():
				customResponse(http.StatusConflict)
			case exceptiongo.TypeOf[types.StatusUnauthorized]():
				// 如果是未登录，返回409
				customResponse(http.StatusUnauthorized)
			case exceptiongo.TypeOf[types.QueryBindingError]():
				customResponse(http.StatusBadRequest)
			case exceptiongo.TypeOf[types.JsonBindingError]():
				customResponse(http.StatusBadRequest)
			case exceptiongo.TypeOf[types.RequestDataInvalidError]():
				customResponse(http.StatusBadRequest)
			case exceptiongo.TypeOf[types.ResourceNotFoundError]():
				customResponse(http.StatusBadRequest)
			// 除以上条件之外，把异常继续向上抛出
			default:
				customResponse(http.StatusInternalServerError)
			}
		}).Deploy()

		c.Next()
	}
}

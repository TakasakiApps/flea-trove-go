package middlewares

import (
	"github.com/TakasakiApps/flea-trove-go/internal/server/engine"
	"github.com/gookit/slog"
	"github.com/hanakogo/digine"
)

var logger *slog.Logger

func Register() {
	logger = digine.Require[slog.Logger](digine.NilLabel)

	// 注册Http异常处理器中间件
	engine.Gin.Use(getHttpExceptionHandler())
	// 注册错误处理中间件，用于包装未知错误
	engine.Gin.Use(getErrorHandler())
	// 注册鉴权中间件
	engine.Gin.Use(getAuthorization([]string{
		"/api/auth/register",
		"/api/auth/login",
		"/api/asset/image/fetch",
		"/api/market/product",
		"/api/market/product/list",
		"/api/market/product/listRandom",
		"/api/market/product/search",
	}...))
}

package server

import (
	"fmt"
	"github.com/TakasakiApps/flea-trove-go/internal/database"
	"github.com/TakasakiApps/flea-trove-go/internal/server/engine"
	"github.com/TakasakiApps/flea-trove-go/internal/server/middlewares"
	"github.com/TakasakiApps/flea-trove-go/internal/server/routers"
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/hanakogo/digine"
	"github.com/hanakogo/exceptiongo"
)

func Run() {
	database.Migrate()

	// 注册所有中间件
	middlewares.Register()
	// 注册所有路由
	routers.Register()

	serverPort := *digine.Require[int](digine.NewLabel("SERVER_PORT"))

	// 启动服务器
	err := engine.Gin.Run(fmt.Sprintf(":%v", serverPort))
	exceptiongo.ThrowErr[types.Fatal](err)
}

package main

import (
	"github.com/TakasakiApps/flea-trove-go/internal/initial"
	"github.com/TakasakiApps/flea-trove-go/internal/server"
	"github.com/hanakogo/exceptiongo"
	"os"
)

func main() {
	// 部署全局异常处理器
	defer exceptiongo.NewExceptionHandler(func(e *exceptiongo.Exception) {
		e.PrintStackTrace()
		os.Exit(1)
	}).Deploy()

	// 初始化相关
	initial.Init()

	// 启动服务器
	server.Run()
}

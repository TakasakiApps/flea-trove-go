package order

import "github.com/gin-gonic/gin"

func RegisterOrderApi(group *gin.RouterGroup) {
	orderGroup := group.Group("/order")

	// 查询订单
	orderGroup.GET("/", GetOrderById)
	// 查询订单列表(可选条件：按用户名查找)
	orderGroup.GET("/list", GetOrderList)

	// 创建订单
	orderGroup.POST("/create", CreateOrder)

	// 更新订单
	orderGroup.PUT("/update", UpdateOrder)
	// 支付订单
	orderGroup.PUT("/pay", PayOrder)
	// 关闭订单
	orderGroup.PUT("/close", CloseOrder)
}

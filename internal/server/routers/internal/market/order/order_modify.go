package order

import (
	"github.com/TakasakiApps/flea-trove-go/internal/database"
	"github.com/TakasakiApps/flea-trove-go/internal/models"
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/TakasakiApps/flea-trove-go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/hanakogo/exceptiongo"
	"github.com/hanakogo/hanakoutilgo"
	"time"
)

func ctxQueryOrderById(c *gin.Context) (order *models.Order) {
	query := utils.CtxBindQuery[struct {
		Id uint `form:"id" binding:"required"`
	}](c)
	order = database.Order().GetOrderById(query.Id)
	if order == nil {
		exceptiongo.ThrowMsgF[types.ResourceNotFoundError]("订单 %v 不存在", query.Id)
	}
	ctxOrderCheck(c, *order)
	return
}

func ctxOrderCheck(c *gin.Context, order models.Order) {
	value, _ := c.Get("user_account")
	if order.User != hanakoutilgo.CastTo[string](value) {
		exceptiongo.ThrowMsgF[types.NotPermittedError]("订单 %v 不是该用户的订单", order.ID)
	}
}

var UpdateOrder gin.HandlerFunc = func(c *gin.Context) {
	order := utils.CtxBindJson[models.Order](c)
	if query := database.Order().GetOrderById(order.ID); query != nil {
		ctxOrderCheck(c, *query)
	}
	effected := database.Order().UpdateOrder(order)
	if effected != 1 {
		exceptiongo.ThrowMsgF[types.ResourceNotFoundError]("更新订单 %v 失败", order.ID)
	}
	utils.CtxRespOK[string](c, "更新订单成功")
}

var PayOrder gin.HandlerFunc = func(c *gin.Context) {
	order := ctxQueryOrderById(c)

	if order.Paid == 1 {
		exceptiongo.ThrowMsgF[types.NotPermittedError]("无法支付已支付的订单")
	}

	if order.Closed == 1 {
		exceptiongo.ThrowMsgF[types.NotPermittedError]("无法支付已关闭的订单")
	}

	// 订单状态设为已支付
	order.Paid = 1
	// 支付成功后自动关闭订单
	order.Closed = 1
	order.UpdatedAt = time.Now()

	effected := database.Order().UpdateOrder(order)
	if effected != 1 {
		exceptiongo.ThrowMsg[types.DBIOError]("支付失败")
	}

	// 完成订单后下架商品，并指定不可再上架
	product := database.Product().GetProductById(order.ProductId)
	product.Published = 0
	product.Sold = 1
	product.CanPublish = 0

	effected = database.Product().UpdateProduct(product)
	if effected != 1 {
		exceptiongo.ThrowMsgF[types.DBIOError]("支付成功，但下架商品失败")
	}

	utils.CtxRespOK[string](c, "支付成功")
}

var CloseOrder gin.HandlerFunc = func(c *gin.Context) {
	order := ctxQueryOrderById(c)

	if order.Closed == 1 {
		exceptiongo.ThrowMsgF[types.NotPermittedError]("无法关闭已关闭的订单")
	}

	// 订单状态设为已关闭
	order.Closed = 1

	effected := database.Order().UpdateOrder(order)
	if effected != 1 {
		exceptiongo.ThrowMsg[types.DBIOError]("订单关闭失败")
	}

	// 订单关闭后，判断订单是否为未支付状态，未支付则重新上架商品，并解除被售出状态
	// 如果已支付，是不可能的，因为支付订单后会自动关闭订单
	if order.Paid == 0 {
		product := database.Product().GetProductById(order.ProductId)
		product.Published = 1
		product.Sold = 0

		effected := database.Product().UpdateProduct(product)
		if effected != 1 {
			exceptiongo.ThrowMsgF[types.DBIOError]("关闭未支付订单成功，但重新上架商品失败")
		}
	}

	utils.CtxRespOK[string](c, "订单关闭成功")
}

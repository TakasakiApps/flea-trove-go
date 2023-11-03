package order

import (
	"github.com/TakasakiApps/flea-trove-go/internal/database"
	"github.com/TakasakiApps/flea-trove-go/internal/models"
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/TakasakiApps/flea-trove-go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/hanakogo/exceptiongo"
	"github.com/hanakogo/hanakoutilgo"
)

var CreateOrder gin.HandlerFunc = func(c *gin.Context) {
	order := utils.CtxBindJson[models.Order](c)

	if queryProduct := database.Product().GetProductById(order.ProductId); queryProduct == nil {
		exceptiongo.ThrowMsgF[types.ResourceNotFoundError]("商品 %v 不存在", order.ProductId)
	}

	// 取出上下文中存储的用户账号
	value, _ := c.Get("user_account")
	userAccount := hanakoutilgo.CastTo[string](value)
	// 把用户账号设置到订单中
	order.User = userAccount

	effected := database.Order().CreateOrder(order)
	if effected != 1 {
		exceptiongo.ThrowMsg[types.DBIOError]("创建失败")
	}

	utils.CtxRespOK[string](c, "创建成功")
}

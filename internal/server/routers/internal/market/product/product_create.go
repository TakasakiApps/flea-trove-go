package product

import (
	"github.com/TakasakiApps/flea-trove-go/internal/database"
	"github.com/TakasakiApps/flea-trove-go/internal/models"
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/TakasakiApps/flea-trove-go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/hanakogo/exceptiongo"
	"github.com/hanakogo/hanakoutilgo"
)

// CreateProduct 创建商品
var CreateProduct gin.HandlerFunc = func(c *gin.Context) {
	product := utils.CtxBindJson[models.Product](c)

	if !utils.ImageExists(product.ImageId) {
		exceptiongo.ThrowMsgF[types.ResourceNotFoundError]("%s 图片不存在", product.ImageId)
	}

	// 取出上下文中存储的用户账号
	value, _ := c.Get("user_account")
	userAccount := hanakoutilgo.CastTo[string](value)
	// 把用户账号设置到商品中
	product.User = userAccount

	effected := database.Product().CreateProduct(product)
	if effected != 1 {
		exceptiongo.ThrowMsg[types.DBIOError]("创建失败")
	}

	utils.CtxRespOK[string](c, "创建成功")
}

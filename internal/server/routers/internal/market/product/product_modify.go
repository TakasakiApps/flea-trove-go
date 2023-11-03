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

func ctxQueryProductById(c *gin.Context) (product *models.Product) {
	query := utils.CtxBindQuery[struct {
		Id uint `form:"id" binding:"required"`
	}](c)
	product = database.Product().GetProductById(query.Id)
	if product == nil {
		exceptiongo.ThrowMsgF[types.ResourceNotFoundError]("商品 %v 不存在", query.Id)
	}
	ctxProductCheck(c, *product)
	return
}

func ctxProductCheck(c *gin.Context, product models.Product) {
	value, _ := c.Get("user_account")
	if product.User != hanakoutilgo.CastTo[string](value) {
		exceptiongo.ThrowMsgF[types.NotPermittedError]("商品 %v 不是该用户的商品", product.ID)
	}
}

// UpdateProduct 更新商品
var UpdateProduct gin.HandlerFunc = func(c *gin.Context) {
	product := utils.CtxBindJson[models.Product](c)
	if query := database.Product().GetProductById(product.ID); query == nil {
		ctxProductCheck(c, *query)
	}
	effected := database.Product().UpdateProduct(product)
	if effected != 1 {
		exceptiongo.ThrowMsg[types.DBIOError]("更新失败")
	}
	utils.CtxRespOK[string](c, "更新成功")
}

// PublishProduct 上架商品
var PublishProduct gin.HandlerFunc = func(c *gin.Context) {
	product := ctxQueryProductById(c)
	product.Published = 1
	effected := database.Product().UpdateProduct(product)
	if effected != 1 {
		exceptiongo.ThrowMsg[types.DBIOError]("上架失败")
	}
	utils.CtxRespOK[string](c, "上架成功")
}

// UnPublishProduct 下架商品
var UnPublishProduct gin.HandlerFunc = func(c *gin.Context) {
	product := ctxQueryProductById(c)
	product.Published = 0
	effected := database.Product().UpdateProduct(product)
	if effected != 1 {
		exceptiongo.ThrowMsg[types.DBIOError]("下架失败")
	}
	utils.CtxRespOK[string](c, "下架成功")
}

// SellProduct 卖出商品
var SellProduct gin.HandlerFunc = func(c *gin.Context) {
	product := ctxQueryProductById(c)
	// 卖出商品自动下架
	product.Published = 0
	product.Sold = 1
	effected := database.Product().UpdateProduct(product)
	if effected != 1 {
		exceptiongo.ThrowMsg[types.DBIOError]("卖出失败")
	}
	utils.CtxRespOK[string](c, "卖出成功")
}

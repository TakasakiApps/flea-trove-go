package product

import (
	"github.com/TakasakiApps/flea-trove-go/internal/database"
	"github.com/TakasakiApps/flea-trove-go/internal/models"
	"github.com/TakasakiApps/flea-trove-go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/mathutil"
)

var GetProductById gin.HandlerFunc = func(c *gin.Context) {
	params := utils.CtxBindQuery[struct {
		Id uint `form:"id" binding:"required"`
	}](c)

	product := database.Product().GetProductById(params.Id)

	utils.CtxRespOK[*models.Product](c, product)
}

var GetProductList gin.HandlerFunc = func(c *gin.Context) {
	params := utils.CtxBindQuery[struct {
		Account string `form:"account"`
	}](c)

	var productList []models.Product
	if params.Account == "" {
		productList = database.Product().GetProductList()
	} else {
		productList = database.Product().GetProductListByUser(params.Account)
	}

	utils.CtxRespOK[[]models.Product](c, productList)
}

var GetProductByKeyword gin.HandlerFunc = func(c *gin.Context) {
	params := utils.CtxBindQuery[struct {
		Keyword string `form:"keyword" binding:"required"`
	}](c)

	productList := database.Product().GetProductListByKeyword(params.Keyword)

	utils.CtxRespOK[[]models.Product](c, productList)
}

var GetProductListRandom gin.HandlerFunc = func(c *gin.Context) {
	params := utils.CtxBindQuery[struct {
		Count int `form:"count"`
	}](c)
	// 默认4个
	if params.Count == 0 {
		params.Count = 4
	}
	// 定义结果切片
	var result []models.Product
	// 定义
	productList := database.Product().GetProductList()

	// 如果数据库里的记录甚至都不够，那就降低需求
	params.Count = mathutil.Min[int](len(productList), params.Count)

	for len(result) < params.Count {
		// 随机获取一个productList长度范围内的下标
		randIndex := mathutil.RandInt(0, len(productList))
		// 取出该下标对应的product
		productRand := productList[randIndex]
		// 放进结果切片里
		result = append(result, productRand)
		// 从productList里删除该product
		productList = append(productList[:randIndex], productList[randIndex+1:]...)
	}

	utils.CtxRespOK[[]models.Product](c, result)
}

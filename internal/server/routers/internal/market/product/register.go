package product

import "github.com/gin-gonic/gin"

func RegisterProductApi(group *gin.RouterGroup) {
	productGroup := group.Group("/product")

	productGroup.GET("/", GetProductById)
	productGroup.GET("/list", GetProductList)
	productGroup.GET("/listRandom", GetProductListRandom)
	productGroup.GET("/search", GetProductByKeyword)

	productGroup.POST("/create", CreateProduct)

	productGroup.PUT("/update", UpdateProduct)
	productGroup.PUT("/publish", PublishProduct)
	productGroup.PUT("/unpublish", UnPublishProduct)
	productGroup.PUT("/sell", SellProduct)
}

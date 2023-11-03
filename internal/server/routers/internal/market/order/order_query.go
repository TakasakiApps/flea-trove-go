package order

import (
	"github.com/TakasakiApps/flea-trove-go/internal/database"
	"github.com/TakasakiApps/flea-trove-go/internal/models"
	"github.com/TakasakiApps/flea-trove-go/internal/utils"
	"github.com/gin-gonic/gin"
)

var GetOrderById gin.HandlerFunc = func(c *gin.Context) {
	params := utils.CtxBindQuery[struct {
		Id uint `form:"id" binding:"required"`
	}](c)

	order := database.Order().GetOrderById(params.Id)

	utils.CtxRespOK[*models.Order](c, order)
}

var GetOrderList gin.HandlerFunc = func(c *gin.Context) {
	params := utils.CtxBindQuery[struct {
		User string `form:"user"`
	}](c)

	var orderList []models.Order
	if params.User == "" {
		orderList = database.Order().GetOrderList()
	} else {
		orderList = database.Order().GetOrderByUser(params.User)
	}

	utils.CtxRespOK[[]models.Order](c, orderList)
}

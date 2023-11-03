package image

import (
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/TakasakiApps/flea-trove-go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/hanakogo/exceptiongo"
)

var FetchImage gin.HandlerFunc = func(c *gin.Context) {
	query := utils.CtxBindQuery[struct {
		Id string `form:"id" binding:"required"`
	}](c)

	if !utils.ImageExists(query.Id) {
		exceptiongo.ThrowMsgF[types.ResourceNotFoundError]("找不到图片 %s", query.Id)
	}

	c.File(utils.ImagePath(query.Id))
}

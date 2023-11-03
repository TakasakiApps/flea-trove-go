package image

import (
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/TakasakiApps/flea-trove-go/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/hanakogo/exceptiongo"
	"net/http"
)

var UploadImage gin.HandlerFunc = func(c *gin.Context) {
	id, err := utils.ImageSave(c)
	exceptiongo.ThrowErr[types.IOError](err)
	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}

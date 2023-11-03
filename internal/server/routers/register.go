package routers

import (
	"github.com/TakasakiApps/flea-trove-go/internal/server/engine"
	"github.com/gin-gonic/gin"
)

var apiGroup *gin.RouterGroup

func Register() {
	apiGroup = engine.Gin.Group("/api")
	registerAuthApi()
	registerAssetApi()
	registerMarketApi()
}

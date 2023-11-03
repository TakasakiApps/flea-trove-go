package routers

import (
	"github.com/TakasakiApps/flea-trove-go/internal/server/routers/internal/asset/image"
)

func registerAssetApi() {
	assetGroup := apiGroup.Group("/asset")
	image.RegisterImageApi(assetGroup)
}

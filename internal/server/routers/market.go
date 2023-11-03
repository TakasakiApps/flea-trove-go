package routers

import (
	"github.com/TakasakiApps/flea-trove-go/internal/server/routers/internal/market/order"
	"github.com/TakasakiApps/flea-trove-go/internal/server/routers/internal/market/product"
)

func registerMarketApi() {
	marketGroup := apiGroup.Group("/market")
	product.RegisterProductApi(marketGroup)
	order.RegisterOrderApi(marketGroup)
}

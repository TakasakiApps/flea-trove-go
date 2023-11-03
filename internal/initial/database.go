package initial

import (
	"github.com/TakasakiApps/flea-trove-go/internal/types"
	"github.com/TakasakiApps/flea-trove-go/internal/utils"
	"github.com/hanakogo/digine"
	"github.com/hanakogo/exceptiongo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDataBase() {
	// 连接SQLITE
	open, err := gorm.Open(sqlite.Open(utils.GetAppDataFile("sqlite.db")), &gorm.Config{
		IgnoreRelationshipsWhenMigrating: true,
	})
	exceptiongo.ThrowErr[types.DBConnectError](err)
	digine.Bind[gorm.DB](open, digine.NilLabel)
}
